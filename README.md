# Akave SDK CLI

The **Akave SDK CLI** (`akavesdk`) is a command-line tool designed to streamline interactions with Akave's decentralized storage platform. With this SDK, developers can efficiently manage buckets, upload and download files, and interact with metadata across Akave nodes. It supports both standard API operations and an advanced streaming API for high-performance file management, as well as an IPC API for blockchain-based data operations.

Whether you're building a new integration or managing data across nodes, this SDK provides robust capabilities to help you achieve seamless, scalable storage solutions.

```Base commit: tag v0.4.4```.

## Build and test instructions
Requirements: Go 1.25+

`make build` - outputs a cli binary into `bin/akavecli`.<br>
`make test` - runs tests.<br>
Look at `Makefile` for details.

### Akave Node IPC API

The Akave Node IPC API provides access to gRPC service that interacts with Akave node, and IPC deployed smart-contract that operates with metadata.
Uses same File, Bucket, Block and Chunk models as regular Akave Node API


| Endpoint             | Description                                                                                                                                                                                                                                                                                                                                                                                                                              |
|----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `ConnectionParams`   | Retrieves dial URI and deployed smart contract address to interract with.                                                                                                                                                                                                                                                                                                                                                                |
| `BucketCreate`       | Unimplemented. Functionality calls from SDK side.                                                                                                                                                                                                                                                                                                                                                                                        |
| `BucketView`         | Retrieves single bucket metadata by bucket name and creator address. Calls smart contract method GetBucketByName, transforms response into API bucket model.                                                                                                                                                                                                                                                                             |
| `BucketList`         | Retrieves all buckets metadata (ID, name, created at). For now doesn't sort by creator address.                                                                                                                                                                                                                                                                                                                                          |
| `BucketDelete`       | Unimplemented. Functionality calls from SDK side.                                                                                                                                                                                                                                                                                                                                                                                        |
| `FileView`           | Retrieves single file metadata by file name, bucket name, creator address. Calls smart contract GetBucketByName, to receive it's ID, then GetFileByName and transforms response into API file model.                                                                                                                                                                                                                                     |
| `FileList`           | Retrieves all files of bucket (by name and creator address) metadata. Calls smart contract GetBucketByName, then GetFileByName through all bucket's file id's list.                                                                                                                                                                                                                                                                      |
| `FileDelete`         | Unimplemented. Functionality calls from SDK side.                                                                                                                                                                                                                                                                                                                                                                                        |
| `FileUploadCreate`   | Initiates a file upload. Selects node for each file block to upload.                                                                                                                                                                                                                                                                                                                                                                     |
| `FileUploadBlock`    | Uploads the given block (block's data) via grpc streaming to the node address. Node stores information about peer ID of a node which now has this block and stores it on smart contract. |
| `FileDownloadCreate` | Fetches the file's metadata and its breakdown on blocks: list of blocks this file is made of from smart contract by calling GetBucketByName, GetFileByName, GetFileBlockById respectively. Assigns peer to each block.                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `FileDownloadBlock`  | Downloads the block by cid via grpc streaming from a node which address picked in **FileDownloadCreate**.                                                                                                                                                                                                                                                                                                                                |

# IPC Storage smart contract

The Storage smart contract contains collection of methods that provides access to bucket, files, chunks and peerblocks metadata.

#### Bucket model
- bucket metadata
    ```go
    type StorageBucket struct {
        Id        [32]byte
        Name      string
        CreatedAt *big.Int
        Owner     common.Address
        Files     [][32]byte
    }
    ```

#### File model
- file metadata
    ```go
    type StorageFile struct {
        Id        [32]byte
        Cid       []byte
        BucketId  [32]byte
        Name      string
        Size      *big.Int
        CreatedAt *big.Int
        Blocks    [][32]byte
    }
    ```

#### Block model
- block metadata
    ```go
    type StorageFileBlock struct {
        Id     [32]byte
        Cid    []byte
        FileId [32]byte
        Size   *big.Int
    }
    ```
  
#### Functions
- CreateBucket(opts *bind.TransactOpts, name string) (*types.Transaction, error) {...} <br>
creates bucket, requires keyed transactor, string name. Used in IPC SDK CreateBucket.
- DeleteBucket(opts *bind.TransactOpts, id [32]byte, name string) (*types.Transaction, error) {...} <br>
deletes bucket, requires same keyed transactor as creator of this bucket, bucket id 32byte array, string name.
Used in IPC SDK DeleteBucket.
- GetBucketByName(opts *bind.CallOpts, name string) (StorageBucket, error) {...} <br>
retrieves bucket metadata, requires From(address of creator) to be filled in request (f.e. &bind.CallOpts{Context: ctx, From: client.Auth.From}), and bucket name.
Used to get bucket's ID in IPC SDK CreateBucket, FileDelete, CreateFileUpload, and in IPC endpoint BucketView, FileView, FileList, FileDownloadCreate.
- AddFile(opts *bind.TransactOpts, cid []byte, bucketId [32]byte, name string, size *big.Int) (*types.Transaction, error) {...} <br>
adds file metadata, requires keyed transactor, content identifier (bytes), bucket id, name and size of *big.Int format, returns transaction.
Used in IPC SDK CreateFileUpload.
- DeleteFile(opts *bind.TransactOpts, id [32]byte, bucketId [32]byte, name string, force bool) (*types.Transaction, error) {...} <br>
deletes file metadata, and if force flag set to true all file blocks. required same keyed transactor as while file add func call, file id, bucket id, name.
Used in IPC SDK FileDelete.
- GetFileByName(opts *bind.CallOpts, bucketId [32]byte, name string) (StorageFile, error) {...} <br>
retrieves file metadata by bucket id, file name.
Used in FileDelete, CreateFileUpload, and IPC endpoint FileDownloadCreate, FileView.
- AddFileBlock(opts *bind.TransactOpts, fileId [32]byte, cid []byte, size *big.Int) (*types.Transaction, error) {...} <br>
adds file block metadata, requires keyed transactor, file id 32byte array, content identifier bytes, size *big.int.
Used in IPC SDK CreateFileUpload.
- GetFileBlockById(opts *bind.CallOpts, id [32]byte) (StorageFileBlock, error) {...} <br>
retrieves file block by id, requires block id as 32byte array.
Used in IPC endpoint FileDownloadCreate and IPC SDK CreateFileDownload.
- AddPeerBlock(opts *bind.TransactOpts, peerId []byte, cid []byte) (*types.Transaction, error) {...} <br>
adds peer block, requires keyed transactor, peerId (node id) bytes, content identifier bytes.
Used in IPC endpoint FileUploadBlock.
- GetPeersByPeerBlockCid(opts *bind.CallOpts, cid []byte) ([][]byte, error) {...} <br>
returns all peerIds that has peer block with given content identifier, requires cid bytes.
Used in IPC endpoint FileDownloadCreate.
<br>

# SDK

### SDK DAG utilities
- `ChunkDAG` a struct that contains node's(in context of a DAG) metainformation: CID, sizes, block metadata
- `DAGRoot` helps to build file's root CID. On each file chunk you have to add a link to chunk(specifiying chunk's CID and its node sizes)
- `BuildDAG` build `ChunkDAG` from a given file. This DAG is of flat structure, meaning it has 1 root and all blocks are children of this single root

### SDK IPC API

| Function Name | Description                                                                                                                                                                                  |
|---------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `CreateBucket` | Creates a new bucket with the specified name calling IPC contract CreateBucket, checks if it exists, returns bucket id.                                                                      |
| `ViewBucket` | Retrieves details of a specific bucket by its name.                                                                                                                                          |
| `ListBuckets` | Lists all buckets stored in IPC contract.                                                                                                                                                    |
| `DeleteBucket` | Soft deletes a specific bucket by its name.                                                                                                                                                  |
| `IPC` | Returns ipc sdk instance that works ipc contract.                                                                                                                                            |
| `ListFiles` | Lists all files from a specific bucket.                                                                                                                                                      |
| `FileInfo` | Retrieves metadata of a specific file by its name and bucket's name.                                                                                                                         |
| `CreateFileUpload` | Initiates a file upload to a specified bucket, adding file's metadata and it's blocks metadata to IPC smart contract. Then Assigns each file blocks to different nodes calling ipc endpoint. |
| `Upload` | Splits file on blocks and performs chunk upload to different nodes. Add peer-block to smart contract, that stores info about where each file block stored node-wise.                         |
| `CreateFileDownload` | Initiates a file download from a specified bucket. Gets a receipt that describes which chunks the file consists of.                                                                          |
| `Download` | Using the receipt returned from `CreateFileDownload` endpoint downloads the file by blocks, previously fetches peer block addresses of blocks                                                |
| `FileDelete` | Hard delete (with all blocks) a specific file by its name and bucket's name.                                                                                                                 |

<br>

# Akave CLI

## Key Features

- **Bucket Management**: Create, view, list, and delete buckets.
- **File Management**: Upload, download, view, list, and delete files.
- **Wallet Management**: Create and manage wallets for Akave storage transactions, import/export private keys and check balances.

Encryption key in flag `-e`(or `--encryption-key` in long form) for upload and download file commands is optional, without using it data will be unencrypted. Key must be a hex encoded string.
The key length **must be 32 bytes** long.

Disable erasure coding flag `disable-erasure-coding` ensures that file is not erasure encoded before being uploaded to akave node(erasure coding is enabled by default).

## Commands

### Bucket Commands
- **Create Bucket**: Creates a new bucket.
  ```sh
  akavecli bucket create <bucket-name> --node-address=connect.akave.ai:5500 --private-key="some-private-key"
  ```
- **Delete Bucket**: Deletes a specific bucket.
  ```sh
  akavecli bucket delete <bucket-name> --node-address=connect.akave.ai:5500 --private-key="some-private-key"
  ```
- **View Bucket**: Retrieves details of a specific bucket.
  ```sh
  akavecli bucket view <bucket-name> --node-address=connect.akave.ai:5500 --private-key="some-private-key"
  ```
- **List Buckets**: List all available buckets.
  ```sh
  akavecli bucket list --node-address=connect.akave.ai:5500 --private-key="some-private-key"
  ```
### File Commands
- **List Files**: List all files in a bucket.
  ```sh
  akavecli file list <bucket-name> --node-address=connect.akave.ai:5500 --private-key="some-private-key"
  ```
- **File Info**: Retrieves file information.
  ```sh
  akavecli file info <bucket-name> <file-name> --node-address=connect.akave.ai:5500 --private-key="some-private-key"
  ```
- **Upload File**: Uploads a file to a bucket(`-e` is optional, key length **must be 32 bytes** long, `disable-erasure-coding` is optional)
  ```sh
  akavecli file upload <bucket-name> <file-path> -e="encryption-key" --disable-erasure-coding --node-address=connect.akave.ai:5500 --private-key="some-private-key"
  ```
- **Download File**: Downloads a file from a bucket(`-e` is optional, key length **must be 32 bytes** long, `disable-erasure-coding` is optional)
  ```sh
  akavecli file download <bucket-name> <file-name> <destination-path> -e="encryption-key" --disable-erasure-coding --node-address=connect.akave.ai:5500 --private-key="some-private-key"
  ```

### Wallet Commands
- **Create Wallet**: Creates a new wallet with a generated private key.
  ```sh
  akavecli wallet create <wallet-name>
  ```
- **List Wallets**: Lists all available wallets with their addresses.
  ```sh
  akavecli wallet list
  ```
- **Export Private Key**: Exports the private key for a specific wallet.
  ```sh
  akavecli wallet export-key <wallet-name>
  ```
- **Import Wallet**: Imports a wallet using an existing private key.
  ```sh
  akavecli wallet import <wallet-name> <private-key>
  ```
- **Check Balance**: Shows the balance for a wallet.
  ```sh
  akavecli wallet balance <wallet-name>
  ```
  Or check the balance of the first available wallet:
  ```sh
  akavecli wallet balance
  ```
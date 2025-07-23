# Security Best Practices

## Private Key Management

### ⚠️ Security Warning

**NEVER** share your private keys or expose them in:
- Source code repositories  
- Configuration files
- Command history
- Environment variables in unsecured environments
- Process lists (avoid inline private keys)

### Recommended Authentication Methods (in order of preference)

#### 1. Wallet Management (Most Secure)

Use the built-in wallet system for secure key storage:

```sh
# Create a new wallet
akavecli wallet create production-wallet

# List available wallets
akavecli wallet list

# Use default wallet (first available)
akavecli ipc bucket create my-bucket --node-address=localhost:5000

# Use specific wallet
akavecli ipc bucket create my-bucket --node-address=localhost:5000 --account=production-wallet
```

**Benefits:**
- Keys are stored encrypted on disk
- No risk of exposure in command history
- Easy to manage multiple accounts
- Built-in address verification

#### 2. Environment Variables (Recommended for CI/CD)

Set private key as environment variable:

```sh
# Unix/Linux/macOS
export AKAVE_PRIVATE_KEY="your-private-key-here"
akavecli ipc bucket create my-bucket --node-address=localhost:5000

# Windows Command Prompt
set AKAVE_PRIVATE_KEY=your-private-key-here
akavecli ipc bucket create my-bucket --node-address=localhost:5000

# Windows PowerShell
$env:AKAVE_PRIVATE_KEY="your-private-key-here"
akavecli ipc bucket create my-bucket --node-address=localhost:5000
```

**Benefits:**
- Not visible in command history
- Suitable for CI/CD pipelines
- Easy to rotate

#### 3. Inline Private Keys (Testing Only)

⚠️ **WARNING: Only use for testing with non-valuable keys**

```sh
akavecli ipc bucket create test-bucket --node-address=localhost:5000 --private-key="test-private-key"
```

**Risks:**
- Visible in command history
- Visible in process lists
- Risk of accidental exposure in logs

## Delete Operations

### Understanding Delete Types

#### Soft Delete (Default for most operations)
- **What it does**: Marks data as deleted but doesn't immediately remove it
- **Recovery**: Data may be recoverable until cleanup cycles run
- **Use case**: Safe deletion with recovery option

```sh
# Soft delete examples
akavecli bucket delete my-bucket --node-address=localhost:5000
akavecli files-streaming delete bucket file.txt --node-address=localhost:5000
```

#### Hard Delete (IPC File Operations)
- **What it does**: Permanently removes data and all associated blocks
- **Recovery**: **Cannot be undone**
- **Use case**: Permanent data removal

```sh
# Hard delete example (IPC only)
akavecli ipc file delete bucket file.txt --node-address=localhost:5000
```

### Best Practices for Delete Operations

1. **Always backup important data** before deletion
2. **Use soft delete** for most operations unless permanent removal is required
3. **Test delete operations** on non-production data first
4. **Verify deletion scope** - understand what data will be affected

## Production Deployment

### Environment Setup

1. **Use dedicated service accounts** with minimal required permissions
2. **Rotate private keys regularly**
3. **Monitor wallet balances** for unexpected changes
4. **Use separate wallets** for different environments (dev/staging/prod)

### CI/CD Integration

```yaml
# Example GitHub Actions workflow
env:
  AKAVE_PRIVATE_KEY: ${{ secrets.AKAVE_PRIVATE_KEY }}
  
steps:
  - name: Deploy to Akave
    run: |
      akavecli ipc file upload production-bucket ./dist/app.tar.gz --node-address=${{ secrets.NODE_ADDRESS }}
```

### Monitoring and Logging

- **Enable audit logging** for all operations
- **Monitor wallet transactions** for unauthorized usage  
- **Set up alerts** for failed authentication attempts
- **Regularly review** access patterns

## Troubleshooting

### Authentication Issues

1. **Verify private key format** (hex-encoded without 0x prefix for wallets)
2. **Check environment variable** is set correctly
3. **Ensure wallet exists** and is accessible
4. **Verify permissions** on wallet files (~/.akave_wallets/)

### Permission Errors

1. **Check wallet file permissions** (should be 600/0600)
2. **Verify account balance** for transaction fees
3. **Ensure correct wallet** is being used for the operation

## Emergency Procedures

### Compromised Private Key

1. **Immediately stop using** the compromised key
2. **Create new wallet** with fresh private key
3. **Transfer assets** to new wallet if possible
4. **Update all systems** to use new credentials
5. **Review logs** for unauthorized transactions

### Lost Private Key

1. **Check backup locations** (if any exist)
2. **Private keys cannot be recovered** - this is by design
3. **Assets associated with lost keys** are permanently inaccessible
4. **Create new wallet** and update systems

## Contact and Support

For security-related issues or questions:
- Review this documentation first
- Check GitHub issues for similar problems
- Contact the development team through official channels

**Remember: Security is a shared responsibility between the platform and users.**

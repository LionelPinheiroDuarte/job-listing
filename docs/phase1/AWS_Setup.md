# AWS Configuration

## Account Information
- User: cli-user
- Region: eu-west-3 (Paris)

## CLI Configuration
- AWS CLI Version: 2.33.7
- Credentials Location: ~/.aws/credentials (NOT COMMITTED)
- Config Location: ~/.aws/config

## Network Information
- My Public IP: [REDACTED - stored locally]
  - Stored in: `~/.devops_ip` (gitignored)

## Security Notes
- Credentials configured: ✅
- Access tested: ✅
- Region optimized for Europe: ✅
- **Credentials files are gitignored and never committed**

## Test Command
```bash
aws sts get-caller-identity
```
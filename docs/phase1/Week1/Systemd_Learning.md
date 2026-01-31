# Systemd Learning Notes

## What is systemd?
Systemd is a server that manage units(services) throught systemctl
a command line tool.

## What is a Service Unit?
A service unit is a service configured by a file that the .service as extension example nginx.service.

## Key Commands
### Service Management
- Start a service: systemctl start [service-name]
- Stop a service: systemctl stop [service-name]
- Restart a service: systemctl restart [service-name]
- Check status: systemctl status [service-name]
### Enable/Disable at Boot
- Enable: systemctl enable [service-name]
- Disable: systemctl disable [service-name]
### View Logs
- View service logs: journalctl -u [service-name]
- Follow logs in real-time: journalctl -f
- View last N lines: journalctl -n [number]

## Service Unit File Structure
[Section Unit]
The unit section: where the service is configured through options like description, documentation and
requires.
[Section Service]
The unit service: Service configuration like working directory, exectStart that start the service and
restart that restart the service after some time.
[Section Install]
The unit install: Information about how to install the unit.
Example with wantedBy that determine the when and how a service start
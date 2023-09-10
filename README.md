# AWS IAM Report Generator

Generate a report on AWS IAM users, including details on creation times, password last used, and access key usages.

## Description

This utility fetches a credential report from AWS IAM and generates a CSV file containing the following columns:
- `username`
- `creation_time`
- `password_last_used`
- `access_key_1_last_used`
- `access_key_2_last_used`

## Prerequisites

Ensure you have the necessary AWS credentials set up in your `~/.aws/credentials` file.

## Usage

1. Clone the repository.
```bash
git clone https://github.com/cloud-craftman-go/aws-iam-reports.git
```

2. Build the project.
```bash
go build -o aws-iam-report cmd/*.go
```

3. Run the tool, specifying the required flags.
```bash
./aws-iam-report -profile=[Your AWS Profile Name] -region=[Desired AWS Region (default: us-east-1)]
```

## Flags

- `-profile`: (Required) Specifies the AWS profile to use from the `~/.aws/credentials` file.
- `-region`: (Optional) Specifies the AWS region where the commands will run. Default is `us-east-1`.

## Output

Upon successful execution, a CSV file will be generated with a name format `aws-report-[profile]-[date].csv` in the current directory.

## Code Structure

- `cmd/main.go`: Entry point of the application.
- `cmd/setup-app.go`: Contains logic to set up the application using command-line flags.
- `cmd/aws-report.go`: Contains the primary logic for generating the AWS report.
- `internal/setup/setup.go`: Defines the primary application structure.

## Contributing

Feel free to fork the repository and submit pull requests. For major changes, please open an issue first.

## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.

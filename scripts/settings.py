import os
import boto3

session = boto3.session.Session()
client = session.client(
    's3',
    region_name='ams3',
    endpoint_url='https://ams3.digitaloceanspaces.com',
    aws_access_key_id=os.getenv('AWS_ACCESS_KEY_ID'),
    aws_secret_access_key=os.getenv('AWS_SECRET_ACCESS_KEY')
)
client.download_file('confetti', 'confetti-dev/confetti-key.pem', './file.pem')

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

# TODO: parameterize arguments
bucket = 'confetti'
source = 'confetti-dev/confetti-dev.yaml'
target = './confetti-dev.yaml'
client.download_file(bucket, source, target)

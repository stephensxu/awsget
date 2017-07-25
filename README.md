# awsget

## Summary

This app is a small CMD tool to use with `aws ec2` [command line interface](http://docs.aws.amazon.com/cli/latest/reference/ec2/)

Almost on daily basis, I needed to ssh into one of our AWS EC2 instances to fix some devops issue.

I'd do the following command in my terminal

ssh deploy@ec2-24-225-111-158.us-west-2.compute.amazonaws.com

But I can never remember or alias the public DNS name of every single machine!
It's very annoying that I'd have to log in my AWS console from browser, do two-step verifications and find out the IP addresses.

I just really, really, really wanted a command that would give a list of instance name with their public DNS name like this from my terminal:

be-1

ec2-24-225-111-158.us-west-2.compute.amazonaws.com

be-2

ec2-25-225-111-158.us-west-2.compute.amazonaws.com

fe-1

ec2-26-225-111-158.us-west-2.compute.amazonaws.com

fe-2

ec2-27-225-111-158.us-west-2.compute.amazonaws.com

stg

ec2-28-225-111-158.us-west-2.compute.amazonaws.com

mdev

ec2-29-225-111-158.us-west-2.compute.amazonaws.com


## Usage

### Installation

Download latest executables for OSX and Linux from [Github releases](https://github.com/stephensxu/awsget/releases).

So far I've only tested and confirmed it's working on darwin_amd64

After downloading put the executable in your `PATH`

`mv ~/Downloads/awsget_darwin_amd64 /usr/local/bin/awsget`

And you can use with `awsget`

If you get permission denied error when trying to execute this, do a chmod 777 on the file

`chmod 777 /usr/local/bin/awsget`

`awsget ssh-ips`

Enjoy!


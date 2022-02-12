# mkpasswd
CLI tool for generating and managing versions of passwords on Google secret manager.

## installation
Download the code to a folder and cd to the folder, then run
```bash
go install
```
Install shell completion. For instance `bash` completion can be installed
by adding following line to your `.bashrc`:
```bash
source <(mkpasswd completion bash)
```

Create a Google cloud project and activate Secrets Manager API. Also
create a service account key and then export following two variables after
replacing the values to your setup:
```bash
GOOGLE_PROJECT_ID=your-project-id
GOOGLE_APPLICATION_CREDENTIALS=service-account-file-path.json
```

## generate passwords
Passwords are named entities that have versions. A new password version
can be generated as follows:
```bash
mkpasswd gen
                  NAME                       PASSWORD       VERSION  
---------------------------------------+------------------+----------
  92c90d2e-6487-4ae1-9c45-457e50b61143   82N=.y`TXk1uvc4l         1  
```
Alternatively, a specific name can be provided
```bash
mkpasswd gen --name=my-passwd
    NAME          PASSWORD       VERSION  
------------+------------------+----------
  my-passwd   IwC9j^hPJ5%t0:v2         1  
```
As you can see, if the named password does not already exist, it's version
will be at `1`. Issuing the same command again using an existing named password
will generate a new version:
```bash
mkpasswd gen --name=my-passwd
    NAME          PASSWORD       VERSION  
------------+------------------+----------
  my-passwd   4Sj7|)q25KRW!PAt         2  
```
```bash
mkpasswd gen --name=my-passwd
    NAME          PASSWORD       VERSION  
------------+------------------+----------
  my-passwd   1%:4BY}E9ouUyqj7         3  
```

Password length, use of symbols, digits and uppercase letters etc. can be
configured.

For instance, an all numeric password can be generated as follows
```bash
mkpasswd gen --length=10 --num-digits=10 --num-symbols=0
                  NAME                    PASSWORD    VERSION  
---------------------------------------+------------+----------
  a826bf6d-4465-4cf2-8630-e7c6bac2f05a   6847301925         1  
```
Similarly, a password with only alphabetical letters can be generated as follows:
```bash
mkpasswd gen --length=10 --num-digits=0 --num-symbols=0
                  NAME                    PASSWORD    VERSION  
---------------------------------------+------------+----------
  dd9f7830-ec3b-4d27-8f4e-cff6b63220a2   SewMLIrobf         1  
```
Use of uppercase letters can be switched off as well:
```bash
mkpasswd gen --length=10 --num-digits=0 --num-symbols=0 --no-uppercase 
                  NAME                    PASSWORD    VERSION  
---------------------------------------+------------+----------
  cebf81b1-6cfa-492f-991f-2fab07e4ed30   mjcofyvrlt         1  
```
And use of repeating characters can be allowed:
```bash
mkpasswd gen --length=20 --num-digits=0 --num-symbols=0 --no-uppercase --allow-repeat 
                  NAME                         PASSWORD         VERSION  
---------------------------------------+----------------------+----------
  188c3b2e-d42e-454f-8e50-40564e058f27   xsobdtkkdqoiwepubodd         1  
```

## retrieve passwords
Stored passwords can be listed:
```bash
mkpasswd list
                  NAME                  
----------------------------------------
  188c3b2e-d42e-454f-8e50-40564e058f27  
  92c90d2e-6487-4ae1-9c45-457e50b61143  
  a826bf6d-4465-4cf2-8630-e7c6bac2f05a  
  cebf81b1-6cfa-492f-991f-2fab07e4ed30  
  dd9f7830-ec3b-4d27-8f4e-cff6b63220a2  
  my-passwd                             
```
And any particular password value can be fetched, which always fetches the
`latest` version of the named password
```bash
mkpasswd get 188c3b2e-d42e-454f-8e50-40564e058f27
                  NAME                         PASSWORD         VERSION  
---------------------------------------+----------------------+----------
  188c3b2e-d42e-454f-8e50-40564e058f27   xsobdtkkdqoiwepubodd         1  
```
A specific version can be fetched using `--version` flag:
```bash
mkpasswd get my-passwd --version=2
    NAME          PASSWORD       VERSION  
------------+------------------+----------
  my-passwd   4Sj7|)q25KRW!PAt         2  
```

## delete password
When a named password is deleted, all versions of secret material are 
deleted forever.
> Please use caution when using this command
```bash
mkpasswd delete 188c3b2e-d42e-454f-8e50-40564e058f27
```

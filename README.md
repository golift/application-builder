Go Lift: Application Builder
===

Go-App Packaging Template for Debian, RedHat, Homebrew and Docker.


The bits are here, and the readme/howto is coming soon.

In a nutshell (and it's a bit more than a nutshell):
-   [Use this template](https://help.github.com/en/articles/creating-a-repository-from-a-template). Name the repo the same as your go app.
-   Edit [settings.sh](https://github.com/golift/application-builder/tree/master/settings.sh). Most of the config is done here.
    -   If you're not sure about something, don't change it.
-   `mv .travis.yml travis.yml ; touch .travis.yml ; travis setup releases`
    -   Accept the defaults. All you need is the `secure` key.
-   Enable Travis-CI builds for your new repo. Just check the box.
-   Copy the `secure` api key from `.travis.yml` to `travis.yml`, then put it back:
    -   Get the indentation right.
    -   `mv travis.yml .travis.yml`
-   With no other changes Travis-CI should build `hello-world` from your repo.

Homebrew
---
If you want to push homebrew formula, you'll need to make another repo for your homebrew tap.
Mine is named `golift/homebrew-mugs`. The `mugs` part can be any word you want, `formula` is normal, I'm just weird.
-   Add an SSH key to your homebrew repo.
    -   Run this: `ssh-keygen -t rsa -b 4096 -C "<your_email>" -f github_deploy_key -N ''`
    -   Make a tar file with this key. `tar -cf .secret_files.tar github_deploy_key` and encrypt it:
    -   `travis encrypt-file .secret_files.tar` and `git add .secret_files.tar.enc`
    -   Select `Add deploy key` and upload the contents of `github_deploy_key.pub` here:
        - `https://github.com/<username>/<repository>/settings/keys`
    -   Update `.travis.yml` with the `$encrypted_XXXXXXXXXXXX_key` and `$encrypted_XXXXXXXXXXXX_iv` values printed by `travis encrypt-file .secret-files.tar`.
    -   Using a tar file in case you want to add more secure files later. Travis only supports one encrypted file per repo.
    -   The included [.travis.yml](.travis.yml) file will extract the tar file and use the github_deploy_key file automatically.

Docker
---
If you want to support Docker Builds, create a new Docker Repo, and link it to your new
Github Repo. Use [this example](https://github.com/golift/application-builder/tree/master/init/docker/hooks) to setup your build in Docker Cloud. Pretty simple, just add the Tag Auto Build and ENV variables.

Linux
---

Supports uploading to packagecloud! This makes providing your packages to Linux users extremely easy.

<a href="https://packagecloud.io"><img src="https://packagecloud.io/images/packagecloud-badge.png" /></a>

To make this work, go to packagecloud.io and copy your API key. Head to terminal and run:

```shell
travis encrypt PACKAGECLOUD-API-KEY
```

That prints out a really long `secure` line.
Copy and paste that into [.travis.yml](.travis.yml) twice, once for each package type (`rpm` and `deb`).

Next
---
1.  Add your Go code, or start hacking on hello-world; it contains quite a bit of boilerplate to get you started.
1.  Update the [MANUAL.md](https://github.com/golift/application-builder/tree/master/examples/MANUAL.md) file.
1.  Refine [settings.sh](https://github.com/golift/application-builder/tree/master/settings.sh) (it doesn't take much). Make some wiki pages, like `Docker`.
1.  If your app runs as a service update the [package scripts for linux](https://github.com/golift/application-builder/tree/master/scripts).
1.  Fix the `REPO` variable in the [install.sh](https://github.com/golift/application-builder/blob/master/scripts/install.sh) script.
1.  Clean up [.gitignore](https://github.com/golift/application-builder/blob/master/.gitignore).

Bitly
---

Bitly API integration is included to allow shortening a download link in the homebrew formula.
Using bitly gives you download statistics for brew tap installs.
There is no other way to obtain this information.
A wiki will later be added explaining how to add your bitly API Token.

Examples
---
These repos use this template.
-   [https://github.com/unifi-poller/unifi-poller](https://github.com/unifi-poller/unifi-poller)
-   [https://github.com/Notifiarr/notifiarr](https://github.com/Notifiarr/notifiarr)
-   [https://github.com/davidnewhall/unpackerr](https://github.com/davidnewhall/unpackerr)
-   [https://github.com/davidnewhall/secspy](https://github.com/davidnewhall/secspy)
-   [https://github.com/golift/turbovanityurls](https://github.com/golift/turbovanityurls)

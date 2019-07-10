Go Lift: Application Builder
===

Go-App Packaging Template for Debian, RedHat, Homebrew and Docker.


The bits are here, and the readme/howto is coming soon.

In a nutshell (and it's a bit more than a nutshell):
-   [Use this template](https://help.github.com/en/articles/creating-a-repository-from-a-template). Name the repo the same as your go app.
-   Edit [.metadata.sh](https://github.com/golift/application-builder/tree/master/.metadata.sh). Most of the config is done there.
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
-   Add an SSH key to your repo. [Follow this](https://github.com/alrra/travis-scripts/blob/master/docs/github-deploy-keys.md), except name the key `.github_deploy_key` and replace the `.github_deploy_key.enc` file in your new repo.
    -   Update `.travis.yml` with the `$encrypted_XXXXXXXXXXXX_key` and `$encrypted_XXXXXXXXXXXX_iv` values printed by `travis encrypt-file .github_deploy_key`.

Docker
---
If you want to support Docker Builds, create a new Docker Repo, and link it to your new
Github Repo. Use [this example](https://github.com/golift/application-builder/tree/master/init/docker/hooks) to setup your build in Docker Cloud. Pretty simple, just add the Tag Auto Build and ENV variables.

Next
---
Add your Go code, or start hacking on hello-world; it contains quite a bit of boilerplate to get you started. Update the [MANUAL.md](https://github.com/golift/application-builder/tree/master/examples/MANUAL.md) file. Refine [.metadata.sh](https://github.com/golift/application-builder/tree/master/.metadata.sh) (it doesn't take much). Make some wiki pages, like `Docker`.
If your app runs as a service update the [package scripts for linux](https://github.com/golift/application-builder/tree/master/scripts). Fix the `REPO` variable in the [install.sh](https://github.com/golift/application-builder/blob/master/scripts/install.sh) script.


Examples
---
These repos use this template.
-   [https://github.com/davidnewhall/unifi-poller](https://github.com/davidnewhall/unifi-poller)
-   [https://github.com/davidnewhall/secspy](https://github.com/davidnewhall/secspy)
-   [https://github.com/davidnewhall/unpacker-poller](https://github.com/davidnewhall/unpacker-poller)

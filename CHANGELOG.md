## 2017.10.17

* Migrated to golang 1.9
* Migrated to packer 1.0.4

## 2017.07.14

* Migrated to packer 1.0.2
* Migrated mitchellh/packer -> hashicorp/packer
* Copying packer code as is, removed glide

## 2017.05.10-1

* Do not render builder's config (issues/#5)
* Pass packer_on_error & packer_user_variables to wrapped builder
* Bump test VM to bento/centos-7.3

## 2017.05.10

* Migrated to golang 1.8
* Migrated to packer 1.0.0
* Strip binaries (~50M -> ~30M)

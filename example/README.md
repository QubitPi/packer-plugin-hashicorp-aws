## The Example Folder

This folder contains a fully working example of the plugin usage. The example defines the `required_plugins` block. A
pre-defined GitHub Action will run `packer init`, `packer validate`, and `packer build` to test the plugin with the
latest version available of Packer.

The folder shall contain multiple HCL2 compatible files. The action will execute Packer at this folder level
running `packer init -upgrade .` and `packer build .`.

If the plugin requires authentication, the configuration should be provided via GitHub Secrets and set as environment
variables in the [test-plugin-example.yml](/.github/workflows/test-plugin-example.yml) file. For example

```yml
  - name: Build
    working-directory: ${{ github.event.inputs.folder }}
    run: PACKER_LOG=${{ github.event.inputs.logs }} packer build .
    env:
      AUTH_KEY: ${{ secrets.AUTH_KEY }}
      AUTH_PASSWORD: ${{ secrets.AUTH_PASSWORD }}
```
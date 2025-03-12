# captive-browser

A more secure, dedicated, Chrome-based captive portal browser that automatically bypasses custom DNS servers.

`captive-browser` detects the DHCP DNS server and runs a SOCKS5 proxy that resolves hostnames through it. Then it starts a Chrome instance in Incognito mode with a separate data directory and waits for it to exit.

This is a fork from [captive-browser](https://github.com/FiloSottile/captive-browser) which was not updated to work with the latest go versions.

## Installation

1. **Download the tar file**:
   Go to the [GitHub releases page](https://github.com/pacoorozco/captive-browser/releases) and download the appropriate `tar.gz` file for your operating system and architecture.

2. **Decompress the tar file**:
   Use the following command to decompress the tar file:
   ```sh
   tar -xzf captive-browser_<OS>_<ARCH>.tar.gz
   ```

3. **Move the binary to a directory in your PATH**:
   ```sh
   sudo mv captive-browser /usr/local/bin/
   ```

4. **Copy the configuration file**:
   You have to install a config file in `$XDG_CONFIG_HOME/captive-browser.toml` (if set) or `~/.config/captive-browser.toml`. You can probably use one of the stock config file from the `config_examples` folder and copy it to your configuration directory:
   ```sh
   cp config_examples/captive-browser-linux-networkmanager.toml ~/.config/captive-browser.toml
   ```

   You can also amend the config to use Chromium instead of Chrome (see `browser` variable).

Replace `<OS>` and `<ARCH>` with your operating system and architecture.

### macOS

```
cp config/captive-browser-mac.toml ~/.config/captive-browser.toml
```

To disable the insecure system captive browser [see here](https://github.com/drduh/macOS-Security-and-Privacy-Guide#captive-portal). If that doesn't work, disable SIP (remember to re-enable it), and rename `/System/Library/CoreServices/Captive Network Assistant.app`.

### Linux
#### Linux with NetworkManager (Ubuntu)

```
cp config/captive-browser-linux-networkmanager.toml ~/.config/captive-browser.toml
```

#### Linux with systemd-networkd

```
cp config/captive-browser-linux-systemd-networkd.toml ~/.config/captive-browser.toml
```

#### Linux with dhcpcd

```
cp config/captive-browser-linux-dhcpd.toml ~/.config/captive-browser.toml
```

## Usage

Simply run `captive-browser`, log into the captive portal, and then *quit* (⌘Q / Ctrl-Q) the Chrome instance.

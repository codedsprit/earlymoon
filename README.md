<h1 align="center"><code>earlymoon</code></h1>
<h4 align="center"> A Dns query tool written in golang 🐟 </h4> </h6>

<h6 align="center">
  <a href="https://github.com/codedsprit/earlymoon#-installation"><code>Installation</code></a>
  *
   <a href="https://github.com/codedsprit/earlymoon#-features"><code>Features</code></a>
  *
  <a href="https://github.com/codedsprit/earlymoon#-contributing"><code>Contribute</code></a>
 </h6>
 <p align="center">
   <a href="https://github.com/codedsprit/earlymoon/releases"><img src="https://img.shields.io/github/v/release/codedsprit/earlymoon?style=flat&amp;labelColor=56534b&amp;color=c1c1b6&amp;logo=GitHub&amp;logoColor=white" alt="GitHub Release"></a>
<a href="https://github.com/codedsprit/earlymoon/issues"><img src="https://img.shields.io/github/issues/codedsprit/earlymoon?style=flat-square&label=Issues&color=white"></a>
<a href="https://github.com/codedsprit/earlymoon/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-white.svg" alt="MIT LICENSE"></a>
<a href="https://ko-fi.com/roshantiwaree"><img src="https://img.shields.io/badge/support-codedsprit%20-white?logo=kofi&logoColor=blue" alt="Ko-fi"></a>


<img src="https://github.com/codedsprit/earlymoon/blob/showcase/earlymoon.png" alt="img" align="right" width="40%" height="20%"></p>
[**`Earlymoon`**](/) is a minimimal yet simple **DNS** query tool. **``earlymoon``** can fetch any dns [records](https://github.com/codedsprit/earlymoon/blob/showcase/supported-records.txt) of given domain, which can be a huge take on **``enumeration``** step.

The project is inspird by ***`dig`*** tool, I have created this tool to improve my skills in  **hacking** & **golang** ecosystem.

<hr>

## Features 🍙
- It's fast.
- Minimal (*only requires **golang** and coreutils*)
- Works on all major **Linux** distrubitions, 


## Installation 📩
    
  <details> <summary><code>Binary </code></summary>
    &nbsp;

  - *Manual* : You can directly download the binary of your arch from [**releases**][releases] and run it.
```
./earlymoon -h
``` 
  </details>

> [!IMPORTANT]
> *_For upstream updates, it's recommended to build `earlymoon` from source !_*

<details open> <summary><code>Source </code></summary>
&nbsp;

   
  ```bash
    git clone --depth=1 --branch main https://github.com/codedsprit/earlymoon
    cd earlymoon
    just build
  ``` 

</details>

<details><summary><code>Dockerfile after clone</code></summary>
&nbsp;

```bash
docker build -t earlymoon .
docker run -it earlymoon
```

</details>
<details> <summary><code>Arch user repository </code></summary>
&nbsp;
  
  ```bash
  paru/yay -S earlymoon-git
  ```

</details>
 <details>
  <summary>Help</summary>
  &nbsp;

  ```bash
~   earlymoon -h

Usage: earlymoon [Flags]
Flags:
    -d       Domain to query (required)
    -t       Type of Dns record (required)
Options:
    -V       Print Version
    help     Display this help message
~ 
```
</details>

## Contributing  🤝
  - Give the project a star ⭐
  - Add new ``commands`` or ``subcommands``
  - Fix code or submit new feature

## Copying 📜
 `earlymoon` is licensed under the **`MIT LICENSE`**. 
 
&nbsp;

<p align="center">
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/footers/gray0_ctp_on_line.svg?sanitize=true" />
</p>
<p align="center">
	Copyright &copy; 2024-present <a href="codedsprit.xyz" target="_blank">codedsprit.xyz</a>
</p>

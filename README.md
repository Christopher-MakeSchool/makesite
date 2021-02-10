# 🔗 makesite

[![Go Report Card](https://goreportcard.com/badge/github.com/Christopher-MakeSchool/makesite)](https://goreportcard.com/report/github.com/Christopher-MakeSchool/makesite)

_Create your own custom Static Site Generator (like [Jekyll](https://jekyllrb.com/) or [Hugo](https://gohugo.io/)) by cloning and fulfilling the requirements in this repo!_

## 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Deliverables](#deliverables)
4. [Resources](#resources)

## Project Structure

```bash
📂 makesite
├── README.md
├── first-post.txt
├── latest-post.txt
├── makesite.go
└── template.tmpl
```

## Getting Started

Frist, fork the [Make-School-Labs/makesite repo](https://github.com/Make-School-Labs/makesite/fork), or [this/my repo](https://github.com/Christopher-MakeSchool/makesite/fork), then run this command

```zsh
git clone git@github.com:YOUR_GITHUB_USERNAME/makesite.git && cd makesite
```

## Deliverables

**For each task**:

- Complete each task in the order they appear.
- Use [GitHub Task List](https://help.github.com/en/github/managing-your-work-on-github/about-task-lists) syntax to update the task list.
- Use [GitHub Version Tags](https://git-scm.com/book/en/v2/Git-Basics-Tagging) to show your progress through this project.

### MVP v1.0 Requirements

Complete the MVP and If you finish early, move on to the stretch challenges. \
If you get stuck on any step, be sure to print the output to `stdout`!

- [x] Edit line `4` of `README.md`. Change this line to the following, replacing `YOUR_USERNAME` and `YOUR_REPONAME` with your GitHub username and repository name respectively.
- [x] Read in the contents of the provided `first-post.txt` file.
- [x] Edit the provided HTML template (`template.tmpl`) to display the contents of `first-post.txt`.
- [x] Render the contents of `first-post.txt` using Go Templates and print it to stdout.
- [X] Write the HTML template to the filesystem to a file. Name it `first-post.html`.
- [X] Manually test the generated HTML page by running `/.makesite`. Double-click the `first-post.html` file that appears in your directory after running the command to open the generated page in your browser.
- [X] **Add, commit, and push to GitHub**.
- [X] Add a new flag to your command named `file`. This flag represents the name of any `.txt` file in the same directory as your program. Run `./makesite --file=latest-post.txt` to test.
- [X] Update the `save` function to use the input filename to generate a new HTML file. For example, if the input file is named `latest-post.txt`, the generated HTML file should be named `latest-post.html`.
- [X] **Add, commit, and push to GitHub**.

#### v1.0 Stretch Challenges

- [X] Use Bootstrap, or another CSS framework, to enhance the style and readability of your template. _Get creative! Writing your very own website generator is a great opportunity to broadcast your style, personality, and development preferences to the world!_

### v1.1

#### v1.1 Requirements

- [X] Create 3 new `.txt` files for testing in the same directory as your project.
- [X] Add a new flag to the `makesite` command named `dir`.
- [X] Use the flag to find all `.txt` files in the given directory. Print them to `stdout`.
- [X] With the list of `.txt` files you found, generate an HTML page for each.
- [X] Run `./makesite --dir=.` to test in your local directory.
- [X] **Add, commit, and push to GitHub**.

#### v1.1 Stretch Challenges

- [X] Recursively find all `.txt` files in the given directory, as well as it's subdirectories. Print them to `stdout` to make sure. Generate an HTML page for each.
- [X] When your program finishes, print: `Success! Generated 5 pages.` The `Success!` substring must be <span style="color: green; font-weight:bold;">bold green</span>, and the count (`5`) must be **bold**.
- [X] Modify the success message to read: `Success! Generated 5 pages (18.2kB total).` Calculate the total by summing the size of each HTML file, then converting the total to kilobytes. Always return one significant digit after the decimal point.
- [ ] Determine how long it took to execute your static site generator. Modify the success message to read: `Success! Generated 5 pages (18.2kB total) in 3.25 seconds.` Always return two significant digits after the decimal point.
- [ ] Test your solutions to these stretch challenges on many different directories containing `.txt` files. Are there any ways to make your code faster?

### v1.2

#### v1.2 Requirements

- [X] Initialize Go modules in your project.
- [X] Add any third party library to your project to enhance it's functionality. Some ideas you might consider include **_(CHOOSE ONLY ONE)_**:
  - [ ] Translating page content using Google Translate.
  - [ ] Parse Markdown (`.md`) files and transform them into HTML. `#` through `######` should translate to `<h1>` through `<h6>` elements.
  - [X] **_FILL IN THE BLANK_**: I will use the `SRG` library. The documentation is located [`here`](https://pkg.go.dev/github.com/foize/go.sgr). My goal is to use it to Complete v1.1 Stretch Challenges of adding color/formating on final print statement.
  - [ ] **_FILL IN THE BLANK_**: I will use the `godotenv` library. The documentation is located [`here`](https://pkg.go.dev/github.com/joho/godotenv@v1.3.0). My goal is to use it to `Hide API Keys & Other Secret Things`.
  - [ ] **_FILL IN THE BLANK_**: I will use the `req` library. The documentation is located [`here`](https://pkg.go.dev/github.com/imroc/req@v0.3.0). My goal is to use it to `Interact with the lichess.org API`.
- [ ] **Add, commit, and push to GitHub**.

## Resources

### Lesson Plans

- [**BEW 2.5**: Project #1 - SSGs](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/SSGProject): Code samples you can use to complete the MVP requirements.
- [**BEW 2.5**: Files & Directories](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/FilesDirectories): Code samples you can use to complete v1.1 requirements.
- [**BEW 2.5**: Files & Directories](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/**3rdPartyLibs**): Code samples you can use to complete v1.2 requirements.

### Example Code

- [**Go By Example**: Reading Files](https://gobyexample.com/reading-files)
- [**Go By Example**: Writing Files](https://gobyexample.com/writing-files)
- [**Go By Example**: Panic](https://gobyexample.com/panic)
- [**GopherAcademy**: Using Go Templates](https://blog.gopheracademy.com/advent-2017/using-go-templates/)
- [**rapid7.com**: Building a Simple CLI Tool with Golang](https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/)

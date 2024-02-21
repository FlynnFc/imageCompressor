# Go Image Processing Tool

## Description

This tool is designed to automate the process of resizing PNG images and converting them to JPEG format, while preserving the directory structure of the input folder. It's built in Go and utilizes concurrency for efficient processing of large numbers of images.

It was made as a replacement for https://squoosh.app/ which I used to manually convert and compress images

## Installation

### Prerequisites

- Go 1.15 or later

### Getting Started

1. Clone the repository to your local machine:

   ```sh
   git clone https://github.com/FlynnFc/imageCompressor.git
   ```

2. Navigate to the cloned directory:

   ```sh
   cd go-image-processing
   ```

3. Build the project:

   ```sh
   go build
   ```

This will create an executable named after your project in the current directory.

## Usage

Run the program from the command line, specifying the required options.

### Options

- `-i`: Path to the input directory containing PNG images.
- `-o`: Path to the output directory where the JPEG images will be saved.
- `-h`: Maximum height of the images (default is 300px).
- `-q`: JPEG quality (1-100, where 100 is the best quality).

### Example

```sh
./go-image-processing -i="path/to/input/folder" -o="path/to/output/folder" -h=300 -q=80
```

This command will process all PNG images found in the specified input directory, ensuring they are no taller than 300 pixels, converting them to JPEG format with a quality of 80, and saving them to the specified output directory while preserving the original subdirectory structure.

## Contributing

We welcome contributions to this project! If you have suggestions for improvements or encounter any issues, please feel free to submit an issue or pull request on GitHub.

- For bugs and feature requests, [open an issue](https://github.com/yourusername/go-image-processing/issues/new).
- For code contributions, please create a pull request with a clear description of your changes.

## License

Specify your project's license here, for example, MIT, GPL, etc.

---

Remember to replace `https://github.com/yourusername/go-image-processing.git` with the actual URL of your GitHub repository and adjust the installation, usage instructions, and contribution guidelines as necessary for your project.

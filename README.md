# Compacto - Media Compression Tool

**Compacto** is a lightweight macOS app (with a CLI for now) that compresses images, videos, and PDFs. The goal of Compacto is to reduce file sizes without sacrificing too much quality, making it easy to manage your media.

## Current features

- **Image Compression**: Supports PNG, JPEG, and GIF.
- **Adjustable Compression Levels**: Set compression quality to suit your needs.
- **Fast and Simple**: Works quickly while maintaining quality.

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/larssonphilip/compacto.git
   cd compacto
   ```

2. **Build the project**:

   ```bash
   go build
   ```

## Usage

### Compress an Image

```bash
./compacto compress-image <inputPath> <outputPath>
```

## Options

- `--quality <1-100>`: Set image compression quality.
- `--speed <1-10>`: Adjust compression speed.
- `--dither <0.0-1.0>`: Set dithering level for images.

## Coming Soon

- **Video Compression**: Compress MP4 videos efficiently.
- **PDF Compression**: Shrink PDF sizes without losing content.
- **Compacto** will soon have a GUI for macOS!

## License

Licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

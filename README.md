# 🚀 gofred Website

A modern, responsive website showcasing the power of [gofred](https://github.com/gofred-io/gofred) - a framework for building web applications using only Go that compiles to WebAssembly.

<div align="center">
<img src="server/img/gofred.png" alt="gofred Logo" width="200" height="200">
</div>

## ✨ What is gofred?

gofred is a revolutionary framework that allows you to build responsive web applications using pure Go code. No JavaScript required - your Go code compiles directly to WebAssembly and runs natively in the browser.

### 🌟 Key Features

- **🎯 Pure Go Development**: Write web applications using only Go
- **⚡ WebAssembly Performance**: Near-native performance in the browser
- **📱 Responsive by Default**: Built-in responsive design system
- **🎨 Modern UI Components**: Rich widget library for beautiful interfaces
- **🔧 Developer Friendly**: Hot reload, comprehensive docs, and great DX
- **🌐 Cross-Platform**: Runs on any modern browser

## 🖥️ Live Demo

Visit the live website: [https://gofred.io](https://gofred.io)

## 🏗️ Architecture

This website demonstrates gofred's capabilities with:

- **Responsive Layout System**: Flexbox-inspired layout widgets
- **Modern Design**: Clean, professional interface
- **Component Architecture**: Reusable UI components
- **State Management**: Reactive state with automatic UI updates
- **Navigation**: Client-side routing with pushstate
- **Documentation**: Comprehensive guides and examples

## 🚀 Quick Start

### Prerequisites

- Go 1.25.1 or higher
- gofred CLI tool
- Docker (for containerized deployment)
- Modern web browser with WebAssembly support

### Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/gofred-io/gofred-website.git
   cd gofred-website
   ```

2. **Install gofred CLI** (if not already installed)
   ```bash
   curl -fsSL https://raw.githubusercontent.com/gofred-io/gofred-cli/refs/heads/master/install.sh | bash
   ```

3. **Run development server**
   ```bash
   gofred app run
   ```

   The development server will automatically open your browser with hot reload enabled on an available port.

### 🐳 Docker Deployment

For production deployment using Docker:

```bash
# Development mode
docker compose up -d
# Access at http://localhost:3000

# Production mode with SSL
docker compose --profile production up -d
# Access at http://localhost and https://localhost
```

**Note**: Docker deployment still uses the containerized approach. For local development, use `gofred app run` instead.

See [DOCKER_DEPLOYMENT.md](DOCKER_DEPLOYMENT.md) for detailed deployment instructions.

## 📚 Project Structure

```
gofred-website/
├── app/                     # Main application code
│   ├── components/          # Reusable UI components
│   │   └── coming_soon/     # Placeholder component
│   ├── pages/               # Page components
│   │   ├── home/            # Homepage sections
│   │   ├── docs/            # Documentation pages
│   │   └── 404/             # Error pages
│   └── theme/               # Global theming
├── Dockerfile               # Production container
├── docker-compose.yml       # Container orchestration
├── nginx.conf               # Web server configuration
└── Makefile                 # Build commands
```

## 🛠️ Development

### Building

```bash
# Run development server with hot reload
gofred app run

# Build for production (coming soon)
gofred app build
```

**Note**: `gofred app build` is currently in development and will be available soon.

### Key Components

- **Home Page** (`app/pages/home/`): Landing page with hero, features, and footer
- **Documentation** (`app/pages/docs/`): Comprehensive guides and API docs
- **Components** (`app/components/`): Reusable UI widgets
- **Theme System** (`app/theme/`): Global styling and breakpoints

### Adding New Pages

1. Create a new page component in `app/pages/`
2. Register the route in the main router
3. Add navigation links as needed
4. Update documentation

## 🎨 Design System

The website uses gofred's built-in design system:

- **Breakpoints**: XS, SM, MD, LG, XL, XXL for responsive design
- **Typography**: Consistent font scaling and hierarchy
- **Colors**: Professional color palette with semantic meanings
- **Spacing**: Consistent spacing scale throughout
- **Components**: Widget-based architecture for reusability

## 📖 Documentation

The website includes comprehensive documentation:

- **Getting Started**: Installation and first app guides
- **Core Concepts**: Widgets, layouts, styling, state management
- **Examples**: Real-world usage patterns
- **API Reference**: Complete widget and method documentation

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for:

- Development setup
- Code style guidelines
- Pull request process
- Issue reporting
- Community guidelines

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- **gofred Framework**: [https://github.com/gofred-io/gofred](https://github.com/gofred-io/gofred)
- **Documentation**: [https://gofred.io/docs](https://gofred.io/docs)
- **Examples**: [https://github.com/gofred-io/examples](https://github.com/gofred-io/examples)
- **Community**: [GitHub Discussions](https://github.com/orgs/gofred-io/discussions)

## 🙏 Acknowledgments

- Go team for WebAssembly support
- Contributors to the gofred framework
- Open source community for inspiration and feedback

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/gofred-io/gofred-website/issues)
- **Discussions**: [GitHub Discussions](https://github.com/orgs/gofred-io/discussions)
- **Documentation**: [gofred.io/docs](https://gofred.io/docs)

---

<div align="center">

**Built with ❤️ using gofred**

[Website](https://gofred.io) • [Documentation](https://gofred.io/docs) • [GitHub](https://github.com/gofred-io/gofred)

</div>

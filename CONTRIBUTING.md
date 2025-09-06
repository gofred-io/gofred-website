# 🤝 Contributing to gofred Website

Thank you for your interest in contributing to the gofred website! This document provides guidelines and information for contributors.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Contributing Guidelines](#contributing-guidelines)
- [Pull Request Process](#pull-request-process)
- [Issue Reporting](#issue-reporting)
- [Code Style](#code-style)
- [Testing](#testing)
- [Documentation](#documentation)

## 🤖 Code of Conduct

This project adheres to a code of conduct adapted from the [Contributor Covenant](https://www.contributor-covenant.org/). By participating, you are expected to uphold this code.

### Our Pledge

We pledge to make participation in our project a harassment-free experience for everyone, regardless of age, body size, disability, ethnicity, gender identity and expression, level of experience, nationality, personal appearance, race, religion, or sexual identity and orientation.

### Our Standards

Examples of behavior that contributes to creating a positive environment include:

- Using welcoming and inclusive language
- Being respectful of differing viewpoints and experiences
- Gracefully accepting constructive criticism
- Focusing on what is best for the community
- Showing empathy towards other community members

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+**: [Download Go](https://golang.org/dl/)
- **Git**: For version control
- **Modern Browser**: Chrome, Firefox, Safari, or Edge with WebAssembly support
- **Docker** (optional): For containerized development

### Fork and Clone

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/gofred-website.git
   cd gofred-website
   ```
3. **Add upstream remote**:
   ```bash
   git remote add upstream https://github.com/gofred-io/gofred-website.git
   ```

## 🛠️ Development Setup

### Local Development

1. **Install dependencies**:
   ```bash
   go mod download
   ```

2. **Start development server**:
   ```bash
   make serve
   # or
   go run server/server.go
   ```

3. **Access the website**:
   ```
   http://localhost:8080
   ```

The development server includes:
- 🔄 **Hot reload**: Automatic rebuilding on Go file changes
- 🌐 **Live server**: Serves static files and WebAssembly
- 📊 **Build status**: Visual feedback on compilation

### Docker Development

```bash
# Start development environment
docker compose up -d

# View logs
docker compose logs -f

# Access at http://localhost:3000
```

### Project Structure

```
gofred-website/
├── app/                    # Application source code
│   ├── components/        # Reusable UI components
│   ├── pages/            # Page components
│   └── theme/            # Global theming
├── server/               # Static assets
│   ├── index.html       # HTML template
│   ├── *.css           # Stylesheets
│   └── img/            # Images
├── docs/               # Documentation
└── scripts/            # Build and utility scripts
```

## 📝 Contributing Guidelines

### Types of Contributions

We welcome various types of contributions:

- 🐛 **Bug fixes**: Fix issues or improve functionality
- ✨ **New features**: Add new pages, components, or functionality
- 📚 **Documentation**: Improve guides, examples, or API docs
- 🎨 **Design**: Enhance UI/UX and visual design
- 🔧 **Infrastructure**: Improve build, deployment, or tooling
- 🧪 **Testing**: Add or improve tests
- 🌍 **Accessibility**: Improve accessibility features

### Contribution Workflow

1. **Check existing issues** to avoid duplicate work
2. **Create an issue** for new features or major changes
3. **Create a branch** from `main` for your work
4. **Make your changes** following our guidelines
5. **Test thoroughly** on multiple browsers
6. **Submit a pull request** with clear description

### Branch Naming

Use descriptive branch names:

- `feature/add-responsive-navigation`
- `fix/mobile-layout-issue`
- `docs/update-installation-guide`
- `refactor/simplify-component-structure`

## 🔄 Pull Request Process

### Before Submitting

- ✅ Code compiles without errors
- ✅ Website works in major browsers (Chrome, Firefox, Safari, Edge)
- ✅ Responsive design works on mobile and desktop
- ✅ No console errors or warnings
- ✅ Documentation updated if needed
- ✅ Follow code style guidelines

### PR Template

When submitting a pull request, include:

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Refactoring
- [ ] Performance improvement

## Testing
- [ ] Tested on Chrome
- [ ] Tested on Firefox
- [ ] Tested on mobile
- [ ] Tested responsive design

## Screenshots (if applicable)
[Include before/after screenshots]

## Related Issues
Fixes #123
```

### Review Process

1. **Automated checks**: CI builds and basic validation
2. **Code review**: Maintainers review code quality and design
3. **Testing**: Manual testing on different browsers/devices
4. **Feedback**: Address review comments and suggestions
5. **Approval**: Merge when approved by maintainers

## 🐛 Issue Reporting

### Bug Reports

Use the bug report template and include:

- **Description**: Clear description of the issue
- **Steps to reproduce**: Detailed steps to recreate the bug
- **Expected behavior**: What should happen
- **Actual behavior**: What actually happens
- **Environment**: Browser, OS, Go version
- **Screenshots**: Visual evidence if applicable

### Feature Requests

Use the feature request template and include:

- **Problem description**: What problem does this solve?
- **Proposed solution**: How should it work?
- **Alternatives considered**: Other approaches you considered
- **Additional context**: Any other relevant information

### Performance Issues

Include:

- **Performance metrics**: Load times, memory usage, etc.
- **Browser developer tools**: Network, performance tabs
- **Device specifications**: Hardware and software details
- **Comparison**: Before/after or expected vs actual

## 🎨 Code Style

### Go Code Style

Follow standard Go conventions:

```go
// Good: Clear, descriptive function names
func CreateResponsiveLayout() widget.Widget {
    return container.New(
        column.New(
            []widget.Widget{
                header(),
                content(),
                footer(),
            },
            column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
        ),
        container.Padding(breakpoint.All(spacing.All(16))),
    )
}

// Good: Proper error handling
func validateInput(input string) error {
    if len(input) == 0 {
        return errors.New("input cannot be empty")
    }
    return nil
}
```

### Component Organization

```go
// component_name.go
package componentname

import (
    // Standard library imports first
    "context"
    "fmt"
    
    // Third-party imports
    "github.com/gofred-io/gofred/foundation/container"
    
    // Local imports
    "github.com/gofred-io/gofred-website/app/theme"
)

// Public component function
func ComponentName(props ComponentProps) widget.Widget {
    return container.New(
        // Component implementation
    )
}

// Private helper functions
func helperFunction() widget.Widget {
    // Helper implementation
}
```

### CSS and Styling

- Use gofred's styling system instead of custom CSS when possible
- Follow responsive design patterns with breakpoints
- Maintain consistent spacing and typography
- Use semantic color names from the theme

### Documentation Style

- Use clear, concise language
- Include code examples for complex concepts
- Add screenshots for visual features
- Keep examples up-to-date with current API

## 🧪 Testing

### Manual Testing Checklist

- [ ] **Functionality**: All features work as expected
- [ ] **Responsive Design**: Layout adapts to different screen sizes
- [ ] **Browser Compatibility**: Works in Chrome, Firefox, Safari, Edge
- [ ] **Performance**: No significant performance regressions
- [ ] **Accessibility**: Keyboard navigation and screen reader friendly
- [ ] **WebAssembly**: WASM loads and executes correctly

### Browser Testing

Test on:
- **Desktop**: Chrome, Firefox, Safari, Edge (latest versions)
- **Mobile**: iOS Safari, Chrome Mobile, Firefox Mobile
- **Screen Sizes**: 320px, 768px, 1024px, 1440px, 1920px

### Performance Testing

- Monitor WebAssembly load times
- Check for memory leaks during navigation
- Verify responsive images load appropriately
- Test on slower network connections

## 📚 Documentation

### Writing Documentation

- **Clear and concise**: Explain concepts simply
- **Code examples**: Include working code samples
- **Visual aids**: Use screenshots and diagrams when helpful
- **Up-to-date**: Keep documentation current with code changes

### Documentation Types

1. **API Documentation**: Function and component references
2. **Tutorials**: Step-by-step guides for common tasks
3. **Examples**: Real-world usage patterns
4. **Architecture**: High-level system design
5. **Deployment**: Setup and deployment guides

### Documentation Standards

- Use Markdown for all documentation
- Include code syntax highlighting
- Add table of contents for long documents
- Cross-reference related documentation
- Include prerequisites and assumptions

## 🚀 Deployment and Release

### Development Workflow

1. **Feature branches**: Develop in feature branches
2. **Pull requests**: Submit PR for review
3. **Code review**: Address feedback and approval
4. **Merge**: Merge to main branch
5. **Deploy**: Automatic deployment to staging
6. **Release**: Manual promotion to production

### Release Process

Releases follow semantic versioning (semver):

- **Major** (1.0.0): Breaking changes
- **Minor** (0.1.0): New features, backwards compatible
- **Patch** (0.0.1): Bug fixes, backwards compatible

## 🌟 Recognition

Contributors will be recognized:

- **Contributors list**: Added to project contributors
- **Release notes**: Mentioned in release announcements
- **Documentation**: Credited in relevant docs
- **Social media**: Highlighted on project social accounts

## 📞 Getting Help

If you need help:

- **Documentation**: Check [gofred.io/docs](https://gofred.io/docs)
- **Discussions**: [GitHub Discussions](https://github.com/orgs/gofred-io/discussions)
- **Issues**: Create an issue for bugs or feature requests
- **Community**: Join our community channels

## 📄 License

By contributing to gofred website, you agree that your contributions will be licensed under the same license as the project (MIT License).

---

Thank you for contributing to gofred! Your efforts help make web development with Go more accessible and enjoyable for everyone. 🚀

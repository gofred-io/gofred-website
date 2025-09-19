package core_concepts

import (
	"github.com/gofred-io/gofred-website/app/components/codeblock"
	appTheme "github.com/gofred-io/gofred-website/app/theme"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func EventHandlingContent() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				eventHandlingPageHeader(),
				spacer.New(spacer.Height(24)),
				eventHandlingPageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func eventHandlingPageHeader() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Event Handling",
				text.FontSize(32),
				text.FontWeight("700"),
			),
			text.New(
				"Learn how to handle user interactions and create responsive, interactive applications with gofred's event system.",
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(18),
			),
		},
		column.Gap(8),
	)
}

func eventHandlingPageContent() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			eventContentSection("Understanding Events", "Events are actions that users perform in your application, such as clicking buttons, typing in inputs, or navigating between pages. gofred provides a simple and powerful event handling system that integrates seamlessly with state management."),
			spacer.New(spacer.Height(24)),

			// Basic Event Handling
			eventContentSection("Basic Event Handling", "Event handlers are functions that respond to user interactions. They receive the widget that triggered the event and an event object containing details about the interaction."),
			spacer.New(spacer.Height(16)),

			eventSubsection("OnClick Events", "The most common event type for handling button clicks and user interactions."),
			codeblock.New(`package main

import (
    "fmt"
    "github.com/gofred-io/gofred/foundation/button"
    "github.com/gofred-io/gofred/foundation/column"
    "github.com/gofred-io/gofred/foundation/text"
    "github.com/gofred-io/gofred/application"
)

// Basic click handler
func simpleButton() application.BaseWidget {
    return button.New(
        text.New("Click Me"),
        button.OnClick(handleSimpleClick),
    )
}

// Event handler function signature
func handleSimpleClick(this application.BaseWidget, e application.Event) {
    fmt.Println("Button was clicked!")
    // 'this' is the widget that was clicked
    // 'e' contains event details
}

// Inline event handler
func inlineHandlerButton() application.BaseWidget {
    return button.New(
        text.New("Inline Handler"),
        button.OnClick(func(this application.BaseWidget, e application.Event) {
            fmt.Println("Inline handler executed!")
            // You can access variables from the surrounding scope
        }),
    )
}`),
			spacer.New(spacer.Height(16)),

			eventSubsection("Event Handler Patterns", "Different patterns for organizing and implementing event handlers."),
			codeblock.New(`// Named handler functions (recommended for reusability)
func saveButton() application.BaseWidget {
    return button.New(
        text.New("Save"),
        button.OnClick(handleSave),
    )
}

func handleSave(this application.BaseWidget, e application.Event) {
    // Dedicated handler function
    fmt.Println("Saving data...")
}

// Closure handlers (for accessing local state)
func counterButton(count int, onIncrement func()) application.BaseWidget {
    return button.New(
        text.New(fmt.Sprintf("Count: %d", count)),
        button.OnClick(func(this application.BaseWidget, e application.Event) {
            // Access closure variables
            onIncrement()
        }),
    )
}

// Method handlers (for widget-based components)
type TodoList struct {
    items []string
}

func (tl *TodoList) addButton() application.BaseWidget {
    return button.New(
        text.New("Add Item"),
        button.OnClick(tl.handleAddItem),
    )
}

func (tl *TodoList) handleAddItem(this application.BaseWidget, e application.Event) {
    tl.items = append(tl.items, "New Item")
    // Update state or trigger re-render
}`),
			spacer.New(spacer.Height(24)),

			// Events with State
			eventContentSection("Events with State Management", "Combine events with state management to create dynamic, reactive user interfaces."),
			spacer.New(spacer.Height(16)),

			eventSubsection("State Updates from Events", "Update application state in response to user interactions."),
			codeblock.New(`import (
    "github.com/gofred-io/gofred/hooks"
    "github.com/gofred-io/gofred/listenable"
)

var (
    count, setCount = hooks.UseState(0)
    message, setMessage = hooks.UseState("Hello")
    isVisible, setIsVisible = hooks.UseState(true)
)

// Counter with state updates
func counterWidget() application.BaseWidget {
    return column.New(
        []application.BaseWidget{
            // Display current count
            listenable.Builder(count, func() application.BaseWidget {
                return text.New(
                    fmt.Sprintf("Count: %d", count.Value()),
                    text.FontSize(18),
                    text.FontWeight("600"),
                )
            }),
            
            // Increment button
            button.New(
                text.New("Increment"),
                button.OnClick(func(this application.BaseWidget, e application.Event) {
                    setCount(count.Value() + 1)
                }),
            ),
            
            // Decrement button
            button.New(
                text.New("Decrement"),
                button.OnClick(func(this application.BaseWidget, e application.Event) {
                    setCount(count.Value() - 1)
                }),
            ),
            
            // Reset button
            button.New(
                text.New("Reset"),
                button.OnClick(func(this application.BaseWidget, e application.Event) {
                    setCount(0)
                }),
            ),
        },
        column.Gap(8),
    )
}`),
			spacer.New(spacer.Height(16)),

			eventSubsection("Conditional Event Handling", "Handle events differently based on current state."),
			codeblock.New(`var (
    isLoggedIn, setIsLoggedIn = hooks.UseState(false)
    user, setUser = hooks.UseState[*User](nil)
)

type User struct {
    Name  string
    Email string
}

func loginButton() application.BaseWidget {
    return listenable.Builder(isLoggedIn, func() application.BaseWidget {
        if isLoggedIn.Value() {
            // Show logout button when logged in
            return button.New(
                text.New("Logout"),
                button.BackgroundColor("#EF4444"),
                button.OnClick(handleLogout),
            )
        } else {
            // Show login button when logged out
            return button.New(
                text.New("Login"),
                button.BackgroundColor("#10B981"),
                button.OnClick(handleLogin),
            )
        }
    })
}

func handleLogin(this application.BaseWidget, e application.Event) {
    // Simulate login process
    fmt.Println("Logging in...")
    
    // Set user data
    setUser(&User{
        Name:  "John Doe",
        Email: "john@example.com",
    })
    
    // Update login state
    setIsLoggedIn(true)
}

func handleLogout(this application.BaseWidget, e application.Event) {
    fmt.Println("Logging out...")
    
    // Clear user data
    setUser(nil)
    setIsLoggedIn(false)
}`),
			spacer.New(spacer.Height(16)),

			eventSubsection("Form Events and Validation", "Handle form submissions and input validation."),
			codeblock.New(`var (
    formData, setFormData = hooks.UseState(map[string]string{
        "email":    "",
        "password": "",
    })
    formErrors, setFormErrors = hooks.UseState(map[string]string{})
    isSubmitting, setIsSubmitting = hooks.UseState(false)
)

func loginForm() application.BaseWidget {
    return column.New(
        []application.BaseWidget{
            // Email field (conceptual - actual input implementation would be here)
            text.New("Email:"),
            
            // Password field (conceptual)
            text.New("Password:"),
            
            // Show validation errors
            listenable.Builder(formErrors, func() application.BaseWidget {
                errors := formErrors.Value()
                if len(errors) == 0 {
                    return spacer.New(spacer.Height(0))
                }
                
                var errorWidgets []application.BaseWidget
                for field, msg := range errors {
                    errorWidgets = append(errorWidgets, text.New(
                        fmt.Sprintf("%s: %s", field, msg),
                        text.FontColor("#EF4444"),
                        text.FontSize(12),
                    ))
                }
                
                return column.New(errorWidgets, column.Gap(4))
            }),
            
            // Submit button
            listenable.Builder(isSubmitting, func() application.BaseWidget {
                if isSubmitting.Value() {
                    return button.New(
                        text.New("Submitting..."),
                        button.Disabled(true),
                    )
                }
                
                return button.New(
                    text.New("Submit"),
                    button.OnClick(handleFormSubmit),
                )
            }),
        },
        column.Gap(12),
    )
}

func handleFormSubmit(this application.BaseWidget, e application.Event) {
    // Prevent double submission
    if isSubmitting.Value() {
        return
    }
    
    // Validate form
    data := formData.Value()
    errors := make(map[string]string)
    
    if data["email"] == "" {
        errors["email"] = "Email is required"
    }
    if data["password"] == "" {
        errors["password"] = "Password is required"
    }
    
    setFormErrors(errors)
    
    if len(errors) > 0 {
        return
    }
    
    // Start submission
    setIsSubmitting(true)
    
    // Simulate async form submission
    go func() {
        // ... submit data to server ...
        
        // Reset form on success
        setFormData(map[string]string{
            "email":    "",
            "password": "",
        })
        setIsSubmitting(false)
    }()
}`),
			spacer.New(spacer.Height(24)),

			// Advanced Event Patterns
			eventContentSection("Advanced Event Patterns", "Learn sophisticated patterns for handling complex user interactions."),
			spacer.New(spacer.Height(16)),

			eventSubsection("Event Delegation and Bubbling", "Handle events from multiple similar widgets efficiently."),
			codeblock.New(`type TodoItem struct {
    ID        string
    Text      string
    Completed bool
}

var (
    todos, setTodos = hooks.UseState([]TodoItem{
        {ID: "1", Text: "Learn gofred", Completed: false},
        {ID: "2", Text: "Build an app", Completed: false},
    })
)

func todoList() application.BaseWidget {
    return listenable.Builder(todos, func() application.BaseWidget {
        items := todos.Value()
        
        var todoWidgets []application.BaseWidget
        for _, item := range items {
            todoWidgets = append(todoWidgets, todoItemWidget(item))
        }
        
        return column.New(todoWidgets, column.Gap(8))
    })
}

func todoItemWidget(item TodoItem) application.BaseWidget {
    return container.New(
        row.New(
            []application.BaseWidget{
                // Toggle completion
                button.New(
                    text.New(func() string {
                        if item.Completed {
                            return "✓"
                        }
                        return "○"
                    }()),
                    button.OnClick(func(this application.BaseWidget, e application.Event) {
                        toggleTodo(item.ID)
                    }),
                ),
                
                // Todo text
                text.New(
                    item.Text,
                    text.FontColor(func() string {
                        if item.Completed {
                            return "#9CA3AF"
                        }
                        return "#1F2937"
                    }()),
                ),
                
                spacer.New(),
                
                // Delete button
                button.New(
                    text.New("Delete"),
                    button.BackgroundColor("#EF4444"),
                    button.OnClick(func(this application.BaseWidget, e application.Event) {
                        deleteTodo(item.ID)
                    }),
                ),
            },
            row.Gap(12),
            row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
        ),
        container.Padding(breakpoint.All(spacing.All(12))),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(spacing.All(1)),
        container.BorderRadius(6),
    )
}

// Event handlers for todo operations
func toggleTodo(id string) {
    items := todos.Value()
    for i, item := range items {
        if item.ID == id {
            items[i].Completed = !items[i].Completed
            break
        }
    }
    setTodos(items)
}

func deleteTodo(id string) {
    items := todos.Value()
    var filtered []TodoItem
    for _, item := range items {
        if item.ID != id {
            filtered = append(filtered, item)
        }
    }
    setTodos(filtered)
}`),
			spacer.New(spacer.Height(16)),

			eventSubsection("Async Event Handling", "Handle events that trigger asynchronous operations."),
			codeblock.New(`var (
    userData, setUserData = hooks.UseState[*User](nil)
    isLoading, setIsLoading = hooks.UseState(false)
    error, setError = hooks.UseState[error](nil)
)

func userProfile() application.BaseWidget {
    return column.New(
        []application.BaseWidget{
            // Load user button
            button.New(
                text.New("Load User Profile"),
                button.OnClick(handleLoadUser),
            ),
            
            // Loading state
            listenable.Builder(isLoading, func() application.BaseWidget {
                if isLoading.Value() {
                    return text.New("Loading...", text.FontColor("#6B7280"))
                }
                return spacer.New(spacer.Height(0))
            }),
            
            // Error state
            listenable.Builder(error, func() application.BaseWidget {
                err := error.Value()
                if err != nil {
                    return column.New(
                        []application.BaseWidget{
                            text.New(
                                fmt.Sprintf("Error: %s", err.Error()),
                                text.FontColor("#EF4444"),
                            ),
                            button.New(
                                text.New("Retry"),
                                button.OnClick(handleLoadUser),
                            ),
                        },
                        column.Gap(8),
                    )
                }
                return spacer.New(spacer.Height(0))
            }),
            
            // User data
            listenable.Builder(userData, func() application.BaseWidget {
                user := userData.Value()
                if user == nil {
                    return text.New("No user data")
                }
                
                return column.New(
                    []application.BaseWidget{
                        text.New(
                            fmt.Sprintf("Name: %s", user.Name),
                            text.FontWeight("600"),
                        ),
                        text.New(
                            fmt.Sprintf("Email: %s", user.Email),
                            text.FontColor("#6B7280"),
                        ),
                        button.New(
                            text.New("Refresh"),
                            button.OnClick(handleLoadUser),
                        ),
                    },
                    column.Gap(8),
                )
            }),
        },
        column.Gap(16),
    )
}

func handleLoadUser(this application.BaseWidget, e application.Event) {
    // Prevent multiple simultaneous requests
    if isLoading.Value() {
        return
    }
    
    // Clear previous error
    setError(nil)
    setIsLoading(true)
    
    // Simulate async API call
    go func() {
        // Simulate network delay
        time.Sleep(2 * time.Second)
        
        // Simulate random success/failure
        if rand.Float32() > 0.7 {
            setError(fmt.Errorf("failed to load user data"))
            setIsLoading(false)
            return
        }
        
        // Success - set user data
        setUserData(&User{
            Name:  "Jane Doe",
            Email: "jane@example.com",
        })
        setIsLoading(false)
    }()
}`),
			spacer.New(spacer.Height(16)),

			eventSubsection("Event Composition and Higher-Order Handlers", "Create reusable event handling patterns."),
			codeblock.New(`// Higher-order function for debounced events
func debounceHandler(delay time.Duration, handler func(application.BaseWidget, application.Event)) func(application.BaseWidget, application.Event) {
    var timer *time.Timer
    
    return func(this application.BaseWidget, e application.Event) {
        if timer != nil {
            timer.Stop()
        }
        
        timer = time.AfterFunc(delay, func() {
            handler(this, e)
        })
    }
}

// Higher-order function for event confirmation
func confirmHandler(message string, handler func(application.BaseWidget, application.Event)) func(application.BaseWidget, application.Event) {
    return func(this application.BaseWidget, e application.Event) {
        // In a real implementation, you'd show a confirmation dialog
        confirmed := true // Simulate user confirmation
        
        if confirmed {
            handler(this, e)
        }
    }
}

// Usage examples
func enhancedButton() application.BaseWidget {
    return column.New(
        []application.BaseWidget{
            // Debounced search button
            button.New(
                text.New("Search"),
                button.OnClick(debounceHandler(300*time.Millisecond, func(this application.BaseWidget, e application.Event) {
                    fmt.Println("Performing search...")
                    // Search logic here
                })),
            ),
            
            // Confirmed delete button
            button.New(
                text.New("Delete All"),
                button.BackgroundColor("#EF4444"),
                button.OnClick(confirmHandler("Are you sure you want to delete all items?", func(this application.BaseWidget, e application.Event) {
                    fmt.Println("Deleting all items...")
                    // Delete logic here
                })),
            ),
        },
        column.Gap(12),
    )
}

// Event handler composition
func compositeHandler(handlers ...func(application.BaseWidget, application.Event)) func(application.BaseWidget, application.Event) {
    return func(this application.BaseWidget, e application.Event) {
        for _, handler := range handlers {
            handler(this, e)
        }
    }
}

// Analytics and logging wrapper
func analyticsHandler(eventName string, handler func(application.BaseWidget, application.Event)) func(application.BaseWidget, application.Event) {
    return func(this application.BaseWidget, e application.Event) {
        // Log analytics event
        fmt.Printf("Analytics: %s triggered\n", eventName)
        
        // Call original handler
        handler(this, e)
    }
}

// Usage with composition
func trackedButton() application.BaseWidget {
    return button.New(
        text.New("Tracked Action"),
        button.OnClick(compositeHandler(
            analyticsHandler("button_clicked", func(this application.BaseWidget, e application.Event) {
                fmt.Println("Primary action executed")
            }),
            func(this application.BaseWidget, e application.Event) {
                fmt.Println("Secondary action executed")
            },
        )),
    )
}`),
			spacer.New(spacer.Height(24)),

			// Navigation Events
			eventContentSection("Navigation and Routing Events", "Handle navigation events and create multi-page applications."),
			spacer.New(spacer.Height(16)),

			eventSubsection("Link Events", "Handle navigation between pages and routes."),
			codeblock.New(`import (
    "github.com/gofred-io/gofred/foundation/link"
)

var (
    currentPage, setCurrentPage = hooks.UseState("home")
    history, setHistory = hooks.UseState([]string{"home"})
)

// Navigation links
func navigationBar() application.BaseWidget {
    return row.New(
        []application.BaseWidget{
            navLink("home", "Home"),
            navLink("about", "About"),
            navLink("contact", "Contact"),
            navLink("docs", "Documentation"),
        },
        row.Gap(16),
    )
}

func navLink(page, label string) application.BaseWidget {
    return listenable.Builder(currentPage, func() application.BaseWidget {
        isActive := currentPage.Value() == page
        
        textColor := "#6B7280"
        if isActive {
            textColor = "#2B799B"
        }
        
        return link.New(
            text.New(label, text.FontColor(textColor)),
            link.Href(fmt.Sprintf("/%s", page)),
            link.OnClick(func(this application.BaseWidget, e application.Event) {
                navigateToPage(page)
            }),
        )
    })
}

func navigateToPage(page string) {
    // Update current page
    setCurrentPage(page)
    
    // Add to history
    hist := history.Value()
    hist = append(hist, page)
    setHistory(hist)
    
    // In a real app, you might also update the browser URL
    fmt.Printf("Navigated to: %s\n", page)
}

// Back button functionality
func backButton() application.BaseWidget {
    return listenable.Builder(history, func() application.BaseWidget {
        hist := history.Value()
        canGoBack := len(hist) > 1
        
        return button.New(
            text.New("← Back"),
            button.Disabled(!canGoBack),
            button.OnClick(func(this application.BaseWidget, e application.Event) {
                if canGoBack {
                    // Remove current page from history
                    newHist := hist[:len(hist)-1]
                    setHistory(newHist)
                    
                    // Go to previous page
                    previousPage := newHist[len(newHist)-1]
                    setCurrentPage(previousPage)
                }
            }),
        )
    })
}`),
			spacer.New(spacer.Height(24)),

			// Best Practices
			eventContentSection("Event Handling Best Practices", "Follow these guidelines for effective event handling."),
			eventHandlingBestPracticesList(),
			spacer.New(spacer.Height(24)),

			eventContentSection("What's Next?", "Now that you understand event handling, explore these related topics:"),
			eventHandlingNextStepsList(),
			spacer.New(spacer.Height(32)),
			navigationButtons("/docs/state", "/docs/examples"),
		},
		column.Gap(16),
	)
}

func eventContentSection(title, description string) application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				title,
				text.FontSize(24),
				text.FontWeight("600"),
			),
			text.New(
				description,
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(16),
			),
		},
		column.Gap(8),
	)
}

func eventSubsection(title, description string) application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				title,
				text.FontSize(20),
				text.FontWeight("600"),
			),
			text.New(
				description,
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(14),
			),
		},
		column.Gap(4),
	)
}

func eventHandlingBestPracticesList() application.BaseWidget {
	practices := []string{
		"Use named functions for event handlers when they can be reused across components",
		"Keep event handlers focused and delegate complex logic to separate functions",
		"Prevent multiple submissions or rapid clicks by checking loading states",
		"Always handle async operations with proper loading and error states",
		"Use closure patterns to access local state and variables in event handlers",
		"Implement proper form validation before processing user input",
		"Consider debouncing events that might fire frequently (like search input)",
		"Use confirmation dialogs for destructive actions like deletions",
		"Log important user interactions for analytics and debugging purposes",
		"Test event handlers thoroughly, including edge cases and error conditions",
		"Clean up resources and timers in event handlers when components unmount",
		"Use higher-order functions to create reusable event handling patterns",
	}

	var practiceItems []application.BaseWidget
	for _, practice := range practices {
		practiceItems = append(practiceItems, eventHandlingListItem(practice))
	}

	return column.New(
		practiceItems,
		column.Gap(8),
	)
}

func eventHandlingListItem(itemText string) application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
			icon.New(
				icondata.Check,
				icon.Width(breakpoint.All(16)),
				icon.Height(breakpoint.All(16)),
				icon.Fill("#10B981"),
			),
			spacer.New(spacer.Width(8)),
			text.New(
				itemText,
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Primary),
				text.FontSize(16),
			),
		},
		row.Gap(8),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func eventHandlingNextStepsList() application.BaseWidget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Examples",
			description: "See complete examples combining widgets, layouts, styling, state, and events",
			href:        "/docs/examples",
		},
	}

	var stepItems []application.BaseWidget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return column.New(
		stepItems,
		column.Gap(12),
	)
}

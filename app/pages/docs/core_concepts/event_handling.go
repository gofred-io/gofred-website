package core_concepts

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/code_block"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/link"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func EventHandlingContent() Widget {
	return Container(
		Column(
			[]Widget{
				eventHandlingPageHeader(),
				Spacer().Height(24),
				eventHandlingPageContent(),
			},
		).Gap(16).Flex(1),
	).Flex(1).Padding(AllBP(All(32)))
}

func eventHandlingPageHeader() Widget {
	return Column(
		[]Widget{
			Text("Event Handling").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Learn how to handle user interactions and create responsive, interactive applications with gofred's event system.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func eventHandlingPageContent() Widget {
	return Column(
		[]Widget{
			eventContentSection("Understanding Events", "Events are actions that users perform in your application, such as clicking buttons, typing in inputs, or navigating between pages. gofred provides a simple and powerful event handling system that integrates seamlessly with state management."),
			Spacer().Height(24),

			// Basic Event Handling
			eventContentSection("Basic Event Handling", "Event handlers are functions that respond to user interactions. They receive the widget that triggered the event and an event object containing details about the interaction."),
			Spacer().Height(16),

			eventSubsection("OnClick Events", "The most common event type for handling button clicks and user interactions."),
			CodeBlock(`package main

import (
    "fmt"
    "github.com/gofred-io/gofred/foundation/button"
    "github.com/gofred-io/gofred/foundation/column"
    "github.com/gofred-io/gofred/foundation/text"
    "github.com/gofred-io/gofred/widget"
)

// Basic click handler
func simpleButton() widget.Widget {
    return button.New(
        text.New("Click Me"),
        button.OnClick(handleSimpleClick),
    )
}

// Event handler function signature
func handleSimpleClick(this widget.Widget, e widget.Event) {
    fmt.Println("Button was clicked!")
    // 'this' is the widget that was clicked
    // 'e' contains event details
}

// Inline event handler
func inlineHandlerButton() widget.Widget {
    return button.New(
        text.New("Inline Handler"),
        button.OnClick(func(this widget.Widget, e widget.Event) {
            fmt.Println("Inline handler executed!")
            // You can access variables from the surrounding scope
        }),
    )
}`),
			Spacer().Height(16),

			eventSubsection("Event Handler Patterns", "Different patterns for organizing and implementing event handlers."),
			CodeBlock(`// Named handler functions (recommended for reusability)
func saveButton() widget.Widget {
    return button.New(
        text.New("Save"),
        button.OnClick(handleSave),
    )
}

func handleSave(this widget.Widget, e widget.Event) {
    // Dedicated handler function
    fmt.Println("Saving data...")
}

// Closure handlers (for accessing local state)
func counterButton(count int, onIncrement func()) widget.Widget {
    return button.New(
        text.New(fmt.Sprintf("Count: %d", count)),
        button.OnClick(func(this widget.Widget, e widget.Event) {
            // Access closure variables
            onIncrement()
        }),
    )
}

// Method handlers (for widget-based components)
type TodoList struct {
    items []string
}

func (tl *TodoList) addButton() widget.Widget {
    return button.New(
        text.New("Add Item"),
        button.OnClick(tl.handleAddItem),
    )
}

func (tl *TodoList) handleAddItem(this widget.Widget, e widget.Event) {
    tl.items = append(tl.items, "New Item")
    // Update state or trigger re-render
}`),
			Spacer().Height(24),

			// Events with State
			eventContentSection("Events with State Management", "Combine events with state management to create dynamic, reactive user interfaces."),
			Spacer().Height(16),

			eventSubsection("State Updates from Events", "Update application state in response to user interactions."),
			CodeBlock(`import (
    "github.com/gofred-io/gofred/hooks"
    "github.com/gofred-io/gofred/listenable"
)

var (
    count, setCount = hooks.UseState(0)
    message, setMessage = hooks.UseState("Hello")
    isVisible, setIsVisible = hooks.UseState(true)
)

// Counter with state updates
func counterWidget() widget.Widget {
    return column.New(
        []widget.Widget{
            // Display current count
            listenable.Builder(count, func() widget.Widget {
                return text.New(
                    fmt.Sprintf("Count: %d", count.Value()),
                    text.FontSize(18),
                    text.FontWeight("600"),
                )
            }),
            
            // Increment button
            button.New(
                text.New("Increment"),
                button.OnClick(func(this widget.Widget, e widget.Event) {
                    setCount(count.Value() + 1)
                }),
            ),
            
            // Decrement button
            button.New(
                text.New("Decrement"),
                button.OnClick(func(this widget.Widget, e widget.Event) {
                    setCount(count.Value() - 1)
                }),
            ),
            
            // Reset button
            button.New(
                text.New("Reset"),
                button.OnClick(func(this widget.Widget, e widget.Event) {
                    setCount(0)
                }),
            ),
        },
        column.Gap(8),
    )
}`),
			Spacer().Height(16),

			eventSubsection("Conditional Event Handling", "Handle events differently based on current state."),
			CodeBlock(`var (
    isLoggedIn, setIsLoggedIn = hooks.UseState(false)
    user, setUser = hooks.UseState[*User](nil)
)

type User struct {
    Name  string
    Email string
}

func loginButton() widget.Widget {
    return listenable.Builder(isLoggedIn, func() widget.Widget {
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

func handleLogin(this widget.Widget, e widget.Event) {
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

func handleLogout(this widget.Widget, e widget.Event) {
    fmt.Println("Logging out...")
    
    // Clear user data
    setUser(nil)
    setIsLoggedIn(false)
}`),
			Spacer().Height(16),

			eventSubsection("Form Events and Validation", "Handle form submissions and input validation."),
			CodeBlock(`var (
    formData, setFormData = hooks.UseState(map[string]string{
        "email":    "",
        "password": "",
    })
    formErrors, setFormErrors = hooks.UseState(map[string]string{})
    isSubmitting, setIsSubmitting = hooks.UseState(false)
)

func loginForm() widget.Widget {
    return column.New(
        []widget.Widget{
            // Email field (conceptual - actual input implementation would be here)
            text.New("Email:"),
            
            // Password field (conceptual)
            text.New("Password:"),
            
            // Show validation errors
            listenable.Builder(formErrors, func() widget.Widget {
                errors := formErrors.Value()
                if len(errors) == 0 {
                    return spacer.New(spacer.Height(0))
                }
                
                var errorWidgets []widget.Widget
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
            listenable.Builder(isSubmitting, func() widget.Widget {
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

func handleFormSubmit(this widget.Widget, e widget.Event) {
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
			Spacer().Height(24),

			// Advanced Event Patterns
			eventContentSection("Advanced Event Patterns", "Learn sophisticated patterns for handling complex user interactions."),
			Spacer().Height(16),

			eventSubsection("Event Delegation and Bubbling", "Handle events from multiple similar widgets efficiently."),
			CodeBlock(`type TodoItem struct {
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

func todoList() widget.Widget {
    return listenable.Builder(todos, func() widget.Widget {
        items := todos.Value()
        
        var todoWidgets []widget.Widget
        for _, item := range items {
            todoWidgets = append(todoWidgets, todoItemWidget(item))
        }
        
        return column.New(todoWidgets, column.Gap(8))
    })
}

func todoItemWidget(item TodoItem) widget.Widget {
    return container.New(
        row.New(
            []widget.Widget{
                // Toggle completion
                button.New(
                    text.New(func() string {
                        if item.Completed {
                            return "✓"
                        }
                        return "○"
                    }()),
                    button.OnClick(func(this widget.Widget, e widget.Event) {
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
                    button.OnClick(func(this widget.Widget, e widget.Event) {
                        deleteTodo(item.ID)
                    }),
                ),
            },
            row.Gap(12),
            row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
        ),
        container.Padding(breakpoint.All(spacing.All(12))),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(1, 1, 1, 1),
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
			Spacer().Height(16),

			eventSubsection("Async Event Handling", "Handle events that trigger asynchronous operations."),
			CodeBlock(`var (
    userData, setUserData = hooks.UseState[*User](nil)
    isLoading, setIsLoading = hooks.UseState(false)
    error, setError = hooks.UseState[error](nil)
)

func userProfile() widget.Widget {
    return column.New(
        []widget.Widget{
            // Load user button
            button.New(
                text.New("Load User Profile"),
                button.OnClick(handleLoadUser),
            ),
            
            // Loading state
            listenable.Builder(isLoading, func() widget.Widget {
                if isLoading.Value() {
                    return text.New("Loading...", text.FontColor("#6B7280"))
                }
                return spacer.New(spacer.Height(0))
            }),
            
            // Error state
            listenable.Builder(error, func() widget.Widget {
                err := error.Value()
                if err != nil {
                    return column.New(
                        []widget.Widget{
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
            listenable.Builder(userData, func() widget.Widget {
                user := userData.Value()
                if user == nil {
                    return text.New("No user data")
                }
                
                return column.New(
                    []widget.Widget{
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

func handleLoadUser(this widget.Widget, e widget.Event) {
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
			Spacer().Height(16),

			eventSubsection("Event Composition and Higher-Order Handlers", "Create reusable event handling patterns."),
			CodeBlock(`// Higher-order function for debounced events
func debounceHandler(delay time.Duration, handler func(widget.Widget, widget.Event)) func(widget.Widget, widget.Event) {
    var timer *time.Timer
    
    return func(this widget.Widget, e widget.Event) {
        if timer != nil {
            timer.Stop()
        }
        
        timer = time.AfterFunc(delay, func() {
            handler(this, e)
        })
    }
}

// Higher-order function for event confirmation
func confirmHandler(message string, handler func(widget.Widget, widget.Event)) func(widget.Widget, widget.Event) {
    return func(this widget.Widget, e widget.Event) {
        // In a real implementation, you'd show a confirmation dialog
        confirmed := true // Simulate user confirmation
        
        if confirmed {
            handler(this, e)
        }
    }
}

// Usage examples
func enhancedButton() widget.Widget {
    return column.New(
        []widget.Widget{
            // Debounced search button
            button.New(
                text.New("Search"),
                button.OnClick(debounceHandler(300*time.Millisecond, func(this widget.Widget, e widget.Event) {
                    fmt.Println("Performing search...")
                    // Search logic here
                })),
            ),
            
            // Confirmed delete button
            button.New(
                text.New("Delete All"),
                button.BackgroundColor("#EF4444"),
                button.OnClick(confirmHandler("Are you sure you want to delete all items?", func(this widget.Widget, e widget.Event) {
                    fmt.Println("Deleting all items...")
                    // Delete logic here
                })),
            ),
        },
        column.Gap(12),
    )
}

// Event handler composition
func compositeHandler(handlers ...func(widget.Widget, widget.Event)) func(widget.Widget, widget.Event) {
    return func(this widget.Widget, e widget.Event) {
        for _, handler := range handlers {
            handler(this, e)
        }
    }
}

// Analytics and logging wrapper
func analyticsHandler(eventName string, handler func(widget.Widget, widget.Event)) func(widget.Widget, widget.Event) {
    return func(this widget.Widget, e widget.Event) {
        // Log analytics event
        fmt.Printf("Analytics: %s triggered\n", eventName)
        
        // Call original handler
        handler(this, e)
    }
}

// Usage with composition
func trackedButton() widget.Widget {
    return button.New(
        text.New("Tracked Action"),
        button.OnClick(compositeHandler(
            analyticsHandler("button_clicked", func(this widget.Widget, e widget.Event) {
                fmt.Println("Primary action executed")
            }),
            func(this widget.Widget, e widget.Event) {
                fmt.Println("Secondary action executed")
            },
        )),
    )
}`),
			Spacer().Height(24),

			// Navigation Events
			eventContentSection("Navigation and Routing Events", "Handle navigation events and create multi-page applications."),
			Spacer().Height(16),

			eventSubsection("Link Events", "Handle navigation between pages and routes."),
			CodeBlock(`import (
    "github.com/gofred-io/gofred/foundation/link"
)

var (
    currentPage, setCurrentPage = hooks.UseState("home")
    history, setHistory = hooks.UseState([]string{"home"})
)

// Navigation links
func navigationBar() widget.Widget {
    return row.New(
        []widget.Widget{
            navLink("home", "Home"),
            navLink("about", "About"),
            navLink("contact", "Contact"),
            navLink("docs", "Documentation"),
        },
        row.Gap(16),
    )
}

func navLink(page, label string) widget.Widget {
    return listenable.Builder(currentPage, func() widget.Widget {
        isActive := currentPage.Value() == page
        
        textColor := "#6B7280"
        if isActive {
            textColor = "#2B799B"
        }
        
        return link.New(
            text.New(label, text.FontColor(textColor)),
            link.Href(fmt.Sprintf("/%s", page)),
            link.OnClick(func(this widget.Widget, e widget.Event) {
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
func backButton() widget.Widget {
    return listenable.Builder(history, func() widget.Widget {
        hist := history.Value()
        canGoBack := len(hist) > 1
        
        return button.New(
            text.New("← Back"),
            button.Disabled(!canGoBack),
            button.OnClick(func(this widget.Widget, e widget.Event) {
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
			Spacer().Height(24),

			// Best Practices
			eventContentSection("Event Handling Best Practices", "Follow these guidelines for effective event handling."),
			eventHandlingBestPracticesList(),
			Spacer().Height(24),

			eventContentSection("What's Next?", "Now that you understand event handling, explore these related topics:"),
			eventHandlingNextStepsList(),
			Spacer().Height(32),
			eventHandlingNavigationButtons("/docs/state", "/docs/examples"),
		},
	).Gap(16)
}

func eventContentSection(title, description string) Widget {
	return Column(
		[]Widget{
			Text(title).
				FontSize(24).
				FontColor("#1F2937").
				FontWeight("600"),
			Text(description).
				FontSize(16).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func eventSubsection(title, description string) Widget {
	return Column(
		[]Widget{
			Text(title).
				FontSize(20).
				FontColor("#1F2937").
				FontWeight("600"),
			Text(description).
				FontSize(14).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(4)
}

func eventHandlingBestPracticesList() Widget {
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

	var practiceItems []Widget
	for _, practice := range practices {
		practiceItems = append(practiceItems, eventHandlingListItem(practice))
	}

	return Column(
		practiceItems,
	).Gap(8)
}

func eventHandlingListItem(itemText string) Widget {
	return Row(
		[]Widget{
			Icon(icondata.Check).
				Width(AllBP(16)).
				Height(AllBP(16)).
				Fill("#10B981"),
			Spacer().Width(8),
			Text(itemText).
				FontSize(16).
				FontColor("#374151").
				FontWeight("400"),
		},
	).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func eventHandlingNextStepsList() Widget {
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

	var stepItems []Widget
	for _, step := range steps {
		stepItems = append(stepItems, eventHandlingNextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}

func eventHandlingNextStepItem(title, description, href string) Widget {
	return Link(
		Container(
			Row(
				[]Widget{
					Column(
						[]Widget{
							Text(title).
								FontSize(16).
								FontColor("#2B799B").
								FontWeight("500"),
							Text(description).
								FontSize(14).
								FontColor("#6B7280").
								FontWeight("400"),
						},
					).Gap(4).Flex(1),
					Icon(icondata.ChevronRight).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#9CA3AF"),
				},
			).Gap(12).Flex(1).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Padding(AllBP(All(16))).
			BackgroundColor("#FFFFFF").
			BorderRadius(8).
			BorderColor("#E5E7EB").
			BorderWidth(1, 1, 1, 1).
			BorderStyle(BorderStyleTypeSolid),
	).Href(href)
}

func eventHandlingNavigationButtons(previousHref, nextHref string) Widget {
	return Row(
		[]Widget{
			Link(
				Button(
					Row(
						[]Widget{
							Icon(icondata.ChevronLeft).
								Width(AllBP(16)).
								Height(AllBP(16)).
								Fill("#FFFFFF"),
							Text("Previous").
								FontSize(14).
								FontColor("#FFFFFF").
								FontWeight("500"),
						},
					).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
				).BackgroundColor("#6B7280").
					Width(AllBP(120)),
			).Href(previousHref),
			Spacer(),
			Link(
				Button(
					Row(
						[]Widget{
							Text("Next").
								FontSize(14).
								FontColor("#FFFFFF").
								FontWeight("500"),
							Icon(icondata.ChevronRight).
								Width(AllBP(16)).
								Height(AllBP(16)).
								Fill("#FFFFFF"),
						},
					).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
				).BackgroundColor("#2B799B").
					Width(AllBP(120)),
			).Href(nextHref),
		},
	).Flex(1)
}

package core_concepts

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	codeblock "github.com/gofred-io/gofred/foundation/code_block"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func StateManagementContent() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				stateManagementPageHeader(),
				spacer.New(spacer.Height(24)),
				stateManagementPageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func stateManagementPageHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"State Management",
				text.FontSize(32),
				text.FontColor("#1F2937"),
				text.FontWeight("700"),
			),
			text.New(
				"Learn how to manage dynamic data and create reactive user interfaces with gofred's powerful state management system.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(8),
	)
}

func stateManagementPageContent() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			stateContentSection("Understanding State", "State represents data that can change over time in your application. In gofred, state management is built around hooks and reactive patterns that automatically update your UI when data changes."),
			spacer.New(spacer.Height(24)),

			// Hooks Introduction
			stateContentSection("Hooks & UseState", "Hooks are functions that let you use state and other gofred features in your widgets. The UseState hook is the foundation of state management."),
			spacer.New(spacer.Height(16)),

			stateSubsection("Basic UseState", "Create stateful values that can be updated and tracked."),
			codeblock.New(`package main

import (
    "fmt"
    "github.com/gofred-io/gofred/hooks"
    "github.com/gofred-io/gofred/listenable"
    "github.com/gofred-io/gofred/foundation/button"
    "github.com/gofred-io/gofred/foundation/column"
    "github.com/gofred-io/gofred/foundation/text"
    "github.com/gofred-io/gofred/widget"
)

var (
    // Declare state variables at package level
    count, setCount = hooks.UseState(0)
    name, setName   = hooks.UseState("World")
)

func counterWidget() widget.BaseWidget {
    return column.New(
        []widget.BaseWidget{
            // Display current count with listenable.Builder
            listenable.Builder(count, func() widget.BaseWidget {
                return text.New(
                    fmt.Sprintf("Count: %d", count.Value()),
                    text.FontSize(18),
                    text.FontWeight("600"),
                )
            }),
            
            // Button to increment count
            button.New(
                text.New("Increment"),
                button.OnClick(func(this widget.BaseWidget, e widget.Event) {
                    setCount(count.Value() + 1)
                }),
            ),
        },
        column.Gap(12),
    )
}`),
			spacer.New(spacer.Height(16)),

			stateSubsection("Multiple State Variables", "Manage multiple pieces of state in your application."),
			codeblock.New(`var (
    // User interface state
    isLoggedIn, setIsLoggedIn = hooks.UseState(false)
    username, setUsername     = hooks.UseState("")
    
    // Application state
    currentPage, setCurrentPage = hooks.UseState("home")
    isLoading, setIsLoading     = hooks.UseState(false)
    
    // Form state
    email, setEmail       = hooks.UseState("")
    password, setPassword = hooks.UseState("")
    errors, setErrors     = hooks.UseState([]string{})
)

// Using multiple state variables
func loginForm() widget.BaseWidget {
    return column.New(
        []widget.BaseWidget{
            // Show loading state
            listenable.Builder(isLoading, func() widget.BaseWidget {
                if isLoading.Value() {
                    return text.New("Logging in...", text.FontColor("#6B7280"))
                }
                return text.New("Please log in", text.FontWeight("500"))
            }),
            
            // Email input (conceptual - actual input would be implemented)
            // emailInput(),
            
            // Login button
            listenable.Builder(email, func() widget.BaseWidget {
                return button.New(
                    text.New("Login"),
                    button.OnClick(handleLogin),
                    // Disable if email is empty
                    button.Disabled(email.Value() == ""),
                )
            }),
        },
        column.Gap(16),
    )
}

func handleLogin(this widget.BaseWidget, e widget.Event) {
    setIsLoading(true)
    // Simulate async login
    go func() {
        // ... login logic ...
        setIsLoading(false)
        setIsLoggedIn(true)
        setUsername(email.Value())
    }()
}`),
			spacer.New(spacer.Height(24)),

			// Listenable and Reactive UI
			stateContentSection("Reactive UI with Listenable", "The listenable.Builder function creates reactive UI components that automatically update when state changes."),
			spacer.New(spacer.Height(16)),

			stateSubsection("Basic Reactive Components", "Build components that respond to state changes."),
			codeblock.New(`// Theme switching example
var (
    isDarkMode, setIsDarkMode = hooks.UseState(false)
)

func themeToggleWidget() widget.BaseWidget {
    return column.New(
        []widget.BaseWidget{
            // Display current theme
            listenable.Builder(isDarkMode, func() widget.BaseWidget {
                theme := "Light Mode"
                if isDarkMode.Value() {
                    theme = "Dark Mode"
                }
                return text.New(
                    fmt.Sprintf("Current theme: %s", theme),
                    text.FontSize(16),
                )
            }),
            
            // Toggle button
            listenable.Builder(isDarkMode, func() widget.BaseWidget {
                label := "Switch to Dark"
                if isDarkMode.Value() {
                    label = "Switch to Light"
                }
                return button.New(
                    text.New(label),
                    button.OnClick(func(this widget.BaseWidget, e widget.Event) {
                        setIsDarkMode(!isDarkMode.Value())
                    }),
                )
            }),
        },
        column.Gap(12),
    )
}`),
			spacer.New(spacer.Height(16)),

			stateSubsection("Conditional Rendering", "Show different UI based on state values."),
			codeblock.New(`var (
    user, setUser         = hooks.UseState[*User](nil)
    isLoggedIn, setLoggedIn = hooks.UseState(false)
)

type User struct {
    Name  string
    Email string
    Role  string
}

func userProfileWidget() widget.BaseWidget {
    return listenable.Builder(isLoggedIn, func() widget.BaseWidget {
        if !isLoggedIn.Value() {
            // Show login prompt
            return container.New(
                column.New(
                    []widget.BaseWidget{
                        text.New("Please log in to continue"),
                        button.New(
                            text.New("Login"),
                            button.OnClick(showLoginForm),
                        ),
                    },
                    column.Gap(12),
                ),
                container.Padding(breakpoint.All(spacing.All(24))),
            )
        }
        
        // Show user profile
        return listenable.Builder(user, func() widget.BaseWidget {
            currentUser := user.Value()
            if currentUser == nil {
                return text.New("Loading user data...")
            }
            
            return container.New(
                column.New(
                    []widget.BaseWidget{
                        text.New(
                            fmt.Sprintf("Welcome, %s!", currentUser.Name),
                            text.FontSize(24),
                            text.FontWeight("600"),
                        ),
                        text.New(
                            fmt.Sprintf("Email: %s", currentUser.Email),
                            text.FontColor("#6B7280"),
                        ),
                        text.New(
                            fmt.Sprintf("Role: %s", currentUser.Role),
                            text.FontColor("#6B7280"),
                        ),
                        button.New(
                            text.New("Logout"),
                            button.OnClick(handleLogout),
                        ),
                    },
                    column.Gap(8),
                ),
                container.Padding(breakpoint.All(spacing.All(24))),
            )
        })
    })
}`),
			spacer.New(spacer.Height(16)),

			stateSubsection("Combining Multiple States", "Create complex reactive UIs by combining multiple state variables."),
			codeblock.New(`var (
    items, setItems       = hooks.UseState([]TodoItem{})
    filter, setFilter     = hooks.UseState("all") // "all", "active", "completed"
    editingId, setEditingId = hooks.UseState("")
)

type TodoItem struct {
    ID          string
    Text        string
    Completed   bool
    CreatedAt   time.Time
}

func todoListWidget() widget.BaseWidget {
    return column.New(
        []widget.BaseWidget{
            // Filter buttons
            listenable.Builder(filter, func() widget.BaseWidget {
                return row.New(
                    []widget.BaseWidget{
                        filterButton("all", "All"),
                        filterButton("active", "Active"),
                        filterButton("completed", "Completed"),
                    },
                    row.Gap(8),
                )
            }),
            
            // Todo items list
            listenable.Builder(items, func() widget.BaseWidget {
                return listenable.Builder(filter, func() widget.BaseWidget {
                    filteredItems := getFilteredItems()
                    
                    var itemWidgets []widget.BaseWidget
                    for _, item := range filteredItems {
                        itemWidgets = append(itemWidgets, todoItemWidget(item))
                    }
                    
                    return column.New(itemWidgets, column.Gap(8))
                })
            }),
            
            // Add new item form
            addTodoForm(),
        },
        column.Gap(16),
    )
}

func filterButton(filterValue, label string) widget.BaseWidget {
    return listenable.Builder(filter, func() widget.BaseWidget {
        isActive := filter.Value() == filterValue
        bgColor := "#F3F4F6"
        textColor := "#6B7280"
        
        if isActive {
            bgColor = "#2B799B"
            textColor = "#FFFFFF"
        }
        
        return button.New(
            text.New(label, text.FontColor(textColor)),
            button.BackgroundColor(bgColor),
            button.OnClick(func(this widget.BaseWidget, e widget.Event) {
                setFilter(filterValue)
            }),
        )
    })
}`),
			spacer.New(spacer.Height(24)),

			// Advanced State Patterns
			stateContentSection("Advanced State Patterns", "Learn advanced patterns for managing complex application state."),
			spacer.New(spacer.Height(16)),

			stateSubsection("State Composition", "Combine related state variables into cohesive patterns."),
			codeblock.New(`// Group related state together
type AppState struct {
    User     *User
    Theme    string
    Language string
    Settings map[string]interface{}
}

var (
    appState, setAppState = hooks.UseState(AppState{
        Theme:    "light",
        Language: "en",
        Settings: make(map[string]interface{}),
    })
)

// Helper functions for updating specific parts of state
func updateUserProfile(newUser *User) {
    currentState := appState.Value()
    currentState.User = newUser
    setAppState(currentState)
}

func updateTheme(newTheme string) {
    currentState := appState.Value()
    currentState.Theme = newTheme
    setAppState(currentState)
}

func updateSetting(key string, value interface{}) {
    currentState := appState.Value()
    currentState.Settings[key] = value
    setAppState(currentState)
}

// Using composed state
func settingsPanel() widget.BaseWidget {
    return listenable.Builder(appState, func() widget.BaseWidget {
        state := appState.Value()
        
        return column.New(
            []widget.BaseWidget{
                // Theme setting
                row.New(
                    []widget.BaseWidget{
                        text.New("Theme:", text.FontWeight("500")),
                        text.New(state.Theme, text.FontColor("#6B7280")),
                        button.New(
                            text.New("Change"),
                            button.OnClick(func(this widget.BaseWidget, e widget.Event) {
                                newTheme := "dark"
                                if state.Theme == "dark" {
                                    newTheme = "light"
                                }
                                updateTheme(newTheme)
                            }),
                        ),
                    },
                    row.Gap(8),
                ),
                
                // Language setting
                row.New(
                    []widget.BaseWidget{
                        text.New("Language:", text.FontWeight("500")),
                        text.New(state.Language, text.FontColor("#6B7280")),
                    },
                    row.Gap(8),
                ),
            },
            column.Gap(12),
        )
    })
}`),
			spacer.New(spacer.Height(16)),

			stateSubsection("Async State Management", "Handle asynchronous operations and loading states."),
			codeblock.New(`type AsyncState[T any] struct {
    Data    T
    Loading bool
    Error   error
}

var (
    userProfile, setUserProfile = hooks.UseState(AsyncState[*User]{
        Loading: false,
        Error:   nil,
    })
    
    todoList, setTodoList = hooks.UseState(AsyncState[[]TodoItem]{
        Data:    []TodoItem{},
        Loading: false,
        Error:   nil,
    })
)

// Helper functions for async operations
func loadUserProfile(userID string) {
    // Set loading state
    setUserProfile(AsyncState[*User]{
        Loading: true,
        Error:   nil,
    })
    
    // Simulate async operation
    go func() {
        user, err := fetchUserFromAPI(userID)
        
        if err != nil {
            setUserProfile(AsyncState[*User]{
                Loading: false,
                Error:   err,
            })
            return
        }
        
        setUserProfile(AsyncState[*User]{
            Data:    user,
            Loading: false,
            Error:   nil,
        })
    }()
}

// UI that handles async state
func userProfileAsyncWidget() widget.BaseWidget {
    return listenable.Builder(userProfile, func() widget.BaseWidget {
        state := userProfile.Value()
        
        if state.Loading {
            return column.New(
                []widget.BaseWidget{
                    text.New("Loading user profile...", text.FontColor("#6B7280")),
                    // You could add a spinner here
                },
                column.Gap(8),
            )
        }
        
        if state.Error != nil {
            return column.New(
                []widget.BaseWidget{
                    text.New(
                        fmt.Sprintf("Error: %s", state.Error.Error()),
                        text.FontColor("#EF4444"),
                    ),
                    button.New(
                        text.New("Retry"),
                        button.OnClick(func(this widget.BaseWidget, e widget.Event) {
                            loadUserProfile("current-user-id")
                        }),
                    ),
                },
                column.Gap(8),
            )
        }
        
        user := state.Data
        if user == nil {
            return text.New("No user data available")
        }
        
        return column.New(
            []widget.BaseWidget{
                text.New(
                    fmt.Sprintf("Welcome, %s!", user.Name),
                    text.FontSize(24),
                    text.FontWeight("600"),
                ),
                text.New(user.Email, text.FontColor("#6B7280")),
            },
            column.Gap(8),
        )
    })
}`),
			spacer.New(spacer.Height(16)),

			stateSubsection("State Validation", "Implement validation and error handling in your state."),
			codeblock.New(`type FormState struct {
    Values map[string]string
    Errors map[string]string
    Touched map[string]bool
    IsValid bool
}

var (
    formState, setFormState = hooks.UseState(FormState{
        Values:  make(map[string]string),
        Errors:  make(map[string]string),
        Touched: make(map[string]bool),
        IsValid: false,
    })
)

// Validation functions
func validateEmail(email string) string {
    if email == "" {
        return "Email is required"
    }
    if !strings.Contains(email, "@") {
        return "Invalid email format"
    }
    return ""
}

func validatePassword(password string) string {
    if password == "" {
        return "Password is required"
    }
    if len(password) < 8 {
        return "Password must be at least 8 characters"
    }
    return ""
}

// Helper functions for form state management
func updateFormValue(field, value string) {
    state := formState.Value()
    state.Values[field] = value
    state.Touched[field] = true
    
    // Validate the field
    var errorMsg string
    switch field {
    case "email":
        errorMsg = validateEmail(value)
    case "password":
        errorMsg = validatePassword(value)
    }
    
    if errorMsg != "" {
        state.Errors[field] = errorMsg
    } else {
        delete(state.Errors, field)
    }
    
    // Check if form is valid
    state.IsValid = len(state.Errors) == 0 && 
                   state.Values["email"] != "" && 
                   state.Values["password"] != ""
    
    setFormState(state)
}

// Form component with validation
func validatedForm() widget.BaseWidget {
    return listenable.Builder(formState, func() widget.BaseWidget {
        state := formState.Value()
        
        return column.New(
            []widget.BaseWidget{
                // Email field
                column.New(
                    []widget.BaseWidget{
                        text.New("Email", text.FontWeight("500")),
                        // Email input would go here
                        text.New(
                            fmt.Sprintf("Current: %s", state.Values["email"]),
                            text.FontSize(12),
                            text.FontColor("#6B7280"),
                        ),
                        // Show error if field is touched and has error
                        func() widget.BaseWidget {
                            if state.Touched["email"] && state.Errors["email"] != "" {
                                return text.New(
                                    state.Errors["email"],
                                    text.FontSize(12),
                                    text.FontColor("#EF4444"),
                                )
                            }
                            return spacer.New(spacer.Height(0))
                        }(),
                    },
                    column.Gap(4),
                ),
                
                // Submit button
                button.New(
                    text.New("Submit"),
                    button.OnClick(func(this widget.BaseWidget, e widget.Event) {
                        if state.IsValid {
                            handleFormSubmit()
                        }
                    }),
                    // Disable if form is not valid
                    button.Disabled(!state.IsValid),
                ),
            },
            column.Gap(16),
        )
    })
}`),
			spacer.New(spacer.Height(24)),

			// Global State Management
			stateContentSection("Global State Management", "Share state across your entire application with global state patterns."),
			spacer.New(spacer.Height(16)),

			stateSubsection("Application Store", "Create a central store for application-wide state."),
			codeblock.New(`// store/app_store.go
package store

import (
    "github.com/gofred-io/gofred/hooks"
    "github.com/gofred-io/gofred/listenable"
)

type AppStore struct {
    User        *User
    Theme       string
    Notifications []Notification
    Router      RouterState
}

type Notification struct {
    ID      string
    Message string
    Type    string // "info", "success", "warning", "error"
}

type RouterState struct {
    CurrentPath string
    History     []string
}

var (
    // Global application store
    appStore, setAppStore = hooks.UseState(AppStore{
        Theme: "light",
        Notifications: []Notification{},
        Router: RouterState{
            CurrentPath: "/",
            History:     []string{"/"},
        },
    })
)

// Store getters
func GetUser() *User {
    return appStore.Value().User
}

func GetTheme() string {
    return appStore.Value().Theme
}

func GetNotifications() []Notification {
    return appStore.Value().Notifications
}

// Store actions
func SetUser(user *User) {
    store := appStore.Value()
    store.User = user
    setAppStore(store)
}

func SetTheme(theme string) {
    store := appStore.Value()
    store.Theme = theme
    setAppStore(store)
}

func AddNotification(notification Notification) {
    store := appStore.Value()
    store.Notifications = append(store.Notifications, notification)
    setAppStore(store)
}

func RemoveNotification(id string) {
    store := appStore.Value()
    var filtered []Notification
    for _, notif := range store.Notifications {
        if notif.ID != id {
            filtered = append(filtered, notif)
        }
    }
    store.Notifications = filtered
    setAppStore(store)
}

// Store listenable for reactive UI
func AppStoreListenable() listenable.Listenable[AppStore] {
    return appStore
}`),
			spacer.New(spacer.Height(16)),

			stateSubsection("Using Global State", "Access and use global state in your components."),
			codeblock.New(`// Using the global store in components
import "your-app/store"

func notificationBar() widget.BaseWidget {
    return listenable.Builder(store.AppStoreListenable(), func() widget.BaseWidget {
        notifications := store.GetNotifications()
        
        if len(notifications) == 0 {
            return spacer.New(spacer.Height(0))
        }
        
        var notifWidgets []widget.BaseWidget
        for _, notif := range notifications {
            notifWidgets = append(notifWidgets, notificationWidget(notif))
        }
        
        return column.New(notifWidgets, column.Gap(8))
    })
}

func notificationWidget(notif store.Notification) widget.BaseWidget {
    var bgColor, textColor string
    switch notif.Type {
    case "success":
        bgColor = "#10B981"
        textColor = "#FFFFFF"
    case "error":
        bgColor = "#EF4444"
        textColor = "#FFFFFF"
    case "warning":
        bgColor = "#F59E0B"
        textColor = "#FFFFFF"
    default:
        bgColor = "#3B82F6"
        textColor = "#FFFFFF"
    }
    
    return container.New(
        row.New(
            []widget.BaseWidget{
                text.New(notif.Message, text.FontColor(textColor)),
                spacer.New(),
                button.New(
                    text.New("×", text.FontColor(textColor)),
                    button.BackgroundColor("transparent"),
                    button.OnClick(func(this widget.BaseWidget, e widget.Event) {
                        store.RemoveNotification(notif.ID)
                    }),
                ),
            },
            row.Gap(8),
            row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
        ),
        container.BackgroundColor(bgColor),
        container.Padding(breakpoint.All(spacing.All(12))),
        container.BorderRadius(6),
    )
}

// Theme-aware component
func themedButton(label string, onClick func(widget.BaseWidget, widget.Event)) widget.BaseWidget {
    return listenable.Builder(store.AppStoreListenable(), func() widget.BaseWidget {
        theme := store.GetTheme()
        
        var bgColor, textColor string
        if theme == "dark" {
            bgColor = "#374151"
            textColor = "#F9FAFB"
        } else {
            bgColor = "#2B799B"
            textColor = "#FFFFFF"
        }
        
        return button.New(
            text.New(label, text.FontColor(textColor)),
            button.BackgroundColor(bgColor),
            button.OnClick(onClick),
        )
    })
}`),
			spacer.New(spacer.Height(24)),

			// Best Practices
			stateContentSection("State Management Best Practices", "Follow these guidelines for effective state management."),
			stateManagementBestPracticesList(),
			spacer.New(spacer.Height(24)),

			stateContentSection("What's Next?", "Now that you understand state management, explore these related topics:"),
			stateManagementNextStepsList(),
			spacer.New(spacer.Height(32)),
			stateManagementNavigationButtons("/docs/styling", "/docs/events"),
		},
		column.Gap(16),
	)
}

func stateContentSection(title, description string) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				title,
				text.FontSize(24),
				text.FontColor("#1F2937"),
				text.FontWeight("600"),
			),
			text.New(
				description,
				text.FontSize(16),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(8),
	)
}

func stateSubsection(title, description string) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				title,
				text.FontSize(20),
				text.FontColor("#1F2937"),
				text.FontWeight("600"),
			),
			text.New(
				description,
				text.FontSize(14),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(4),
	)
}

func stateManagementBestPracticesList() widget.BaseWidget {
	practices := []string{
		"Declare state variables at package level for global access and persistence",
		"Use descriptive names for state variables and their setters (e.g., user, setUser)",
		"Keep state immutable - create new objects rather than modifying existing ones",
		"Use listenable.Builder to create reactive UI components that respond to state changes",
		"Group related state variables together for better organization",
		"Handle loading and error states explicitly in async operations",
		"Validate state changes and provide user feedback for invalid inputs",
		"Use helper functions to encapsulate complex state update logic",
		"Keep state updates simple and predictable - avoid side effects in setters",
		"Consider using global state stores for application-wide data",
		"Test state changes thoroughly, especially async operations and error conditions",
		"Document your state structure and update patterns for team collaboration",
	}

	var practiceItems []widget.BaseWidget
	for _, practice := range practices {
		practiceItems = append(practiceItems, stateManagementListItem(practice))
	}

	return column.New(
		practiceItems,
		column.Gap(8),
	)
}

func stateManagementListItem(itemText string) widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			icon.New(
				icondata.Check,
				icon.Width(breakpoint.All(16)),
				icon.Height(breakpoint.All(16)),
				icon.Fill("#10B981"),
			),
			spacer.New(spacer.Width(8)),
			text.New(
				itemText,
				text.FontSize(16),
				text.FontColor("#374151"),
				text.FontWeight("400"),
			),
		},
		row.Gap(8),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

func stateManagementNextStepsList() widget.BaseWidget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Events",
			description: "Learn how to handle user interactions and create event-driven applications",
			href:        "/docs/events",
		},
	}

	var stepItems []widget.BaseWidget
	for _, step := range steps {
		stepItems = append(stepItems, stateManagementNextStepItem(step.title, step.description, step.href))
	}

	return column.New(
		stepItems,
		column.Gap(12),
	)
}

func stateManagementNextStepItem(title, description, href string) widget.BaseWidget {
	return link.New(
		container.New(
			row.New(
				[]widget.BaseWidget{
					column.New(
						[]widget.BaseWidget{
							text.New(
								title,
								text.FontSize(16),
								text.FontColor("#2B799B"),
								text.FontWeight("500"),
							),
							text.New(
								description,
								text.FontSize(14),
								text.FontColor("#6B7280"),
								text.FontWeight("400"),
							),
						},
						column.Gap(4),
						column.Flex(1),
					),
					icon.New(
						icondata.ChevronRight,
						icon.Width(breakpoint.All(20)),
						icon.Height(breakpoint.All(20)),
						icon.Fill("#9CA3AF"),
					),
				},
				row.Gap(12),
				row.Flex(1),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			container.Padding(breakpoint.All(spacing.All(16))),
			container.BackgroundColor("#FFFFFF"),
			container.BorderRadius(8),
			container.BorderColor("#E5E7EB"),
			container.BorderWidth(1, 1, 1, 1),
			container.BorderStyle(options.BorderStyleTypeSolid),
		),
		link.Href(href),
	)
}

func stateManagementNavigationButtons(previousHref, nextHref string) widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			link.New(
				button.New(
					row.New(
						[]widget.BaseWidget{
							icon.New(
								icondata.ChevronLeft,
								icon.Width(breakpoint.All(16)),
								icon.Height(breakpoint.All(16)),
								icon.Fill("#FFFFFF"),
							),
							text.New(
								"Previous",
								text.FontSize(14),
								text.FontColor("#FFFFFF"),
								text.FontWeight("500"),
							),
						},
						row.Gap(8),
						row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					),
					button.BackgroundColor("#6B7280"),
					button.Width(breakpoint.All(120)),
				),
				link.Href(previousHref),
			),
			spacer.New(),
			link.New(
				button.New(
					row.New(
						[]widget.BaseWidget{
							text.New(
								"Next",
								text.FontSize(14),
								text.FontColor("#FFFFFF"),
								text.FontWeight("500"),
							),
							icon.New(
								icondata.ChevronRight,
								icon.Width(breakpoint.All(16)),
								icon.Height(breakpoint.All(16)),
								icon.Fill("#FFFFFF"),
							),
						},
						row.Gap(8),
						row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					),
					button.BackgroundColor("#2B799B"),
					button.Width(breakpoint.All(120)),
				),
				link.Href(nextHref),
			),
		},
		row.Flex(1),
	)
}

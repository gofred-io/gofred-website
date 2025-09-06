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

func StateManagementContent() Widget {
	return Container(
		Column(
			[]Widget{
				stateManagementPageHeader(),
				Spacer().Height(24),
				stateManagementPageContent(),
			},
		).Gap(16).Flex(1),
	).Flex(1).Padding(AllBP(All(32)))
}

func stateManagementPageHeader() Widget {
	return Column(
		[]Widget{
			Text("State Management").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Learn how to manage dynamic data and create reactive user interfaces with gofred's powerful state management system.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func stateManagementPageContent() Widget {
	return Column(
		[]Widget{
			stateContentSection("Understanding State", "State represents data that can change over time in your application. In gofred, state management is built around hooks and reactive patterns that automatically update your UI when data changes."),
			Spacer().Height(24),

			// Hooks Introduction
			stateContentSection("Hooks & UseState", "Hooks are functions that let you use state and other gofred features in your widgets. The UseState hook is the foundation of state management."),
			Spacer().Height(16),

			stateSubsection("Basic UseState", "Create stateful values that can be updated and tracked."),
			CodeBlock(`package main

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

func counterWidget() Widget {
    return Column(
        []Widget{
            // Display current count with listenable.Builder
            listenable.Builder(count, func() Widget {
                return Text(
                    fmt.Sprintf("Count: %d", count.Value()),
                    .FontSize(18),
                    .FontWeight("600"),
                )
            }),
            
            // Button to increment count
            Button(
                Text("Increment"),
                button.OnClick(func(this Widget, e widget.Event) {
                    setCount(count.Value() + 1)
                }),
            ),
        },
        .Gap(12),
    )
}`),
			Spacer().Height(16),

			stateSubsection("Multiple State Variables", "Manage multiple pieces of state in your application."),
			CodeBlock(`var (
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
func loginForm() Widget {
    return Column(
        []Widget{
            // Show loading state
            listenable.Builder(isLoading, func() Widget {
                if isLoading.Value() {
                    return Text("Logging in...", .FontColor("#6B7280"))
                }
                return Text("Please log in", .FontWeight("500"))
            }),
            
            // Email input (conceptual - actual input would be implemented)
            // emailInput(),
            
            // Login button
            listenable.Builder(email, func() Widget {
                return Button(
                    Text("Login"),
                    button.OnClick(handleLogin),
                    // Disable if email is empty
                    button.Disabled(email.Value() == ""),
                )
            }),
        },
        .Gap(16),
    )
}

func handleLogin(this Widget, e widget.Event) {
    setIsLoading(true)
    // Simulate async login
    go func() {
        // ... login logic ...
        setIsLoading(false)
        setIsLoggedIn(true)
        setUsername(email.Value())
    }()
}`),
			Spacer().Height(24),

			// Listenable and Reactive UI
			stateContentSection("Reactive UI with Listenable", "The listenable.Builder function creates reactive UI components that automatically update when state changes."),
			Spacer().Height(16),

			stateSubsection("Basic Reactive Components", "Build components that respond to state changes."),
			CodeBlock(`// Theme switching example
var (
    isDarkMode, setIsDarkMode = hooks.UseState(false)
)

func themeToggleWidget() Widget {
    return Column(
        []Widget{
            // Display current theme
            listenable.Builder(isDarkMode, func() Widget {
                theme := "Light Mode"
                if isDarkMode.Value() {
                    theme = "Dark Mode"
                }
                return Text(
                    fmt.Sprintf("Current theme: %s", theme),
                    .FontSize(16),
                )
            }),
            
            // Toggle button
            listenable.Builder(isDarkMode, func() Widget {
                label := "Switch to Dark"
                if isDarkMode.Value() {
                    label = "Switch to Light"
                }
                return Button(
                    Text(label),
                    button.OnClick(func(this Widget, e widget.Event) {
                        setIsDarkMode(!isDarkMode.Value())
                    }),
                )
            }),
        },
        .Gap(12),
    )
}`),
			Spacer().Height(16),

			stateSubsection("Conditional Rendering", "Show different UI based on state values."),
			CodeBlock(`var (
    user, setUser         = hooks.UseState[*User](nil)
    isLoggedIn, setLoggedIn = hooks.UseState(false)
)

type User struct {
    Name  string
    Email string
    Role  string
}

func userProfileWidget() Widget {
    return listenable.Builder(isLoggedIn, func() Widget {
        if !isLoggedIn.Value() {
            // Show login prompt
            return Container(
                Column(
                    []Widget{
                        Text("Please log in to continue"),
                        Button(
                            Text("Login"),
                            button.OnClick(showLoginForm),
                        ),
                    },
                    .Gap(12),
                ),
                .Padding(AllBP(All(24))),
            )
        }
        
        // Show user profile
        return listenable.Builder(user, func() Widget {
            currentUser := user.Value()
            if currentUser == nil {
                return Text("Loading user data...")
            }
            
            return Container(
                Column(
                    []Widget{
                        Text(
                            fmt.Sprintf("Welcome, %s!", currentUser.Name),
                            .FontSize(24),
                            .FontWeight("600"),
                        ),
                        Text(
                            fmt.Sprintf("Email: %s", currentUser.Email),
                            .FontColor("#6B7280"),
                        ),
                        Text(
                            fmt.Sprintf("Role: %s", currentUser.Role),
                            .FontColor("#6B7280"),
                        ),
                        Button(
                            Text("Logout"),
                            button.OnClick(handleLogout),
                        ),
                    },
                    .Gap(8),
                ),
                .Padding(AllBP(All(24))),
            )
        })
    })
}`),
			Spacer().Height(16),

			stateSubsection("Combining Multiple States", "Create complex reactive UIs by combining multiple state variables."),
			CodeBlock(`var (
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

func todoListWidget() Widget {
    return Column(
        []Widget{
            // Filter buttons
            listenable.Builder(filter, func() Widget {
                return Row(
                    []Widget{
                        filterButton("all", "All"),
                        filterButton("active", "Active"),
                        filterButton("completed", "Completed"),
                    },
                    .Gap(8),
                )
            }),
            
            // Todo items list
            listenable.Builder(items, func() Widget {
                return listenable.Builder(filter, func() Widget {
                    filteredItems := getFilteredItems()
                    
                    var itemWidgets []Widget
                    for _, item := range filteredItems {
                        itemWidgets = append(itemWidgets, todoItemWidget(item))
                    }
                    
                    return Column(itemWidgets, .Gap(8))
                })
            }),
            
            // Add new item form
            addTodoForm(),
        },
        .Gap(16),
    )
}

func filterButton(filterValue, label string) Widget {
    return listenable.Builder(filter, func() Widget {
        isActive := filter.Value() == filterValue
        bgColor := "#F3F4F6"
        textColor := "#6B7280"
        
        if isActive {
            bgColor = "#2B799B"
            textColor = "#FFFFFF"
        }
        
        return Button(
            Text(label, .FontColor(textColor)),
            .BackgroundColor(bgColor),
            button.OnClick(func(this Widget, e widget.Event) {
                setFilter(filterValue)
            }),
        )
    })
}`),
			Spacer().Height(24),

			// Advanced State Patterns
			stateContentSection("Advanced State Patterns", "Learn advanced patterns for managing complex application state."),
			Spacer().Height(16),

			stateSubsection("State Composition", "Combine related state variables into cohesive patterns."),
			CodeBlock(`// Group related state together
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
func settingsPanel() Widget {
    return listenable.Builder(appState, func() Widget {
        state := appState.Value()
        
        return Column(
            []Widget{
                // Theme setting
                Row(
                    []Widget{
                        Text("Theme:", .FontWeight("500")),
                        Text(state.Theme, .FontColor("#6B7280")),
                        Button(
                            Text("Change"),
                            button.OnClick(func(this Widget, e widget.Event) {
                                newTheme := "dark"
                                if state.Theme == "dark" {
                                    newTheme = "light"
                                }
                                updateTheme(newTheme)
                            }),
                        ),
                    },
                    .Gap(8),
                ),
                
                // Language setting
                Row(
                    []Widget{
                        Text("Language:", .FontWeight("500")),
                        Text(state.Language, .FontColor("#6B7280")),
                    },
                    .Gap(8),
                ),
            },
            .Gap(12),
        )
    })
}`),
			Spacer().Height(16),

			stateSubsection("Async State Management", "Handle asynchronous operations and loading states."),
			CodeBlock(`type AsyncState[T any] struct {
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
func userProfileAsyncWidget() Widget {
    return listenable.Builder(userProfile, func() Widget {
        state := userProfile.Value()
        
        if state.Loading {
            return Column(
                []Widget{
                    Text("Loading user profile...", .FontColor("#6B7280")),
                    // You could add a spinner here
                },
                .Gap(8),
            )
        }
        
        if state.Error != nil {
            return Column(
                []Widget{
                    Text(
                        fmt.Sprintf("Error: %s", state.Error.Error()),
                        .FontColor("#EF4444"),
                    ),
                    Button(
                        Text("Retry"),
                        button.OnClick(func(this Widget, e widget.Event) {
                            loadUserProfile("current-user-id")
                        }),
                    ),
                },
                .Gap(8),
            )
        }
        
        user := state.Data
        if user == nil {
            return Text("No user data available")
        }
        
        return Column(
            []Widget{
                Text(
                    fmt.Sprintf("Welcome, %s!", user.Name),
                    .FontSize(24),
                    .FontWeight("600"),
                ),
                Text(user.Email, .FontColor("#6B7280")),
            },
            .Gap(8),
        )
    })
}`),
			Spacer().Height(16),

			stateSubsection("State Validation", "Implement validation and error handling in your state."),
			CodeBlock(`type FormState struct {
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
func validatedForm() Widget {
    return listenable.Builder(formState, func() Widget {
        state := formState.Value()
        
        return Column(
            []Widget{
                // Email field
                Column(
                    []Widget{
                        Text("Email", .FontWeight("500")),
                        // Email input would go here
                        Text(
                            fmt.Sprintf("Current: %s", state.Values["email"]),
                            .FontSize(12),
                            .FontColor("#6B7280"),
                        ),
                        // Show error if field is touched and has error
                        func() Widget {
                            if state.Touched["email"] && state.Errors["email"] != "" {
                                return Text(
                                    state.Errors["email"],
                                    .FontSize(12),
                                    .FontColor("#EF4444"),
                                )
                            }
                            return Spacer(.Height(0))
                        }(),
                    },
                    .Gap(4),
                ),
                
                // Submit button
                Button(
                    Text("Submit"),
                    button.OnClick(func(this Widget, e widget.Event) {
                        if state.IsValid {
                            handleFormSubmit()
                        }
                    }),
                    // Disable if form is not valid
                    button.Disabled(!state.IsValid),
                ),
            },
            .Gap(16),
        )
    })
}`),
			Spacer().Height(24),

			// Global State Management
			stateContentSection("Global State Management", "Share state across your entire application with global state patterns."),
			Spacer().Height(16),

			stateSubsection("Application Store", "Create a central store for application-wide state."),
			CodeBlock(`// store/app_store.go
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
			Spacer().Height(16),

			stateSubsection("Using Global State", "Access and use global state in your components."),
			CodeBlock(`// Using the global store in components
import "your-app/store"

func notificationBar() Widget {
    return listenable.Builder(store.AppStoreListenable(), func() Widget {
        notifications := store.GetNotifications()
        
        if len(notifications) == 0 {
            return Spacer(.Height(0))
        }
        
        var notifWidgets []Widget
        for _, notif := range notifications {
            notifWidgets = append(notifWidgets, notificationWidget(notif))
        }
        
        return Column(notifWidgets, .Gap(8))
    })
}

func notificationWidget(notif store.Notification) Widget {
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
    
    return Container(
        Row(
            []Widget{
                Text(notif.Message, .FontColor(textColor)),
                Spacer(),
                Button(
                    Text("×", .FontColor(textColor)),
                    .BackgroundColor("transparent"),
                    button.OnClick(func(this Widget, e widget.Event) {
                        store.RemoveNotification(notif.ID)
                    }),
                ),
            },
            .Gap(8),
            .CrossAxisAlignment(AxisAlignmentTypeCenter),
        ),
        container.BackgroundColor(bgColor),
        .Padding(AllBP(All(12))),
        container.BorderRadius(6),
    )
}

// Theme-aware component
func themedButton(label string, onClick func(Widget, widget.Event)) Widget {
    return listenable.Builder(store.AppStoreListenable(), func() Widget {
        theme := store.GetTheme()
        
        var bgColor, textColor string
        if theme == "dark" {
            bgColor = "#374151"
            textColor = "#F9FAFB"
        } else {
            bgColor = "#2B799B"
            textColor = "#FFFFFF"
        }
        
        return Button(
            Text(label, .FontColor(textColor)),
            .BackgroundColor(bgColor),
            button.OnClick(onClick),
        )
    })
}`),
			Spacer().Height(24),

			// Best Practices
			stateContentSection("State Management Best Practices", "Follow these guidelines for effective state management."),
			stateManagementBestPracticesList(),
			Spacer().Height(24),

			stateContentSection("What's Next?", "Now that you understand state management, explore these related topics:"),
			stateManagementNextStepsList(),
			Spacer().Height(32),
			stateManagementNavigationButtons("/docs/styling", "/docs/events"),
		},
	).Gap(16)
}

func stateContentSection(title, description string) Widget {
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

func stateSubsection(title, description string) Widget {
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

func stateManagementBestPracticesList() Widget {
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

	var practiceItems []Widget
	for _, practice := range practices {
		practiceItems = append(practiceItems, stateManagementListItem(practice))
	}

	return Column(
		practiceItems,
	).Gap(8)
}

func stateManagementListItem(itemText string) Widget {
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

func stateManagementNextStepsList() Widget {
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

	var stepItems []Widget
	for _, step := range steps {
		stepItems = append(stepItems, stateManagementNextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}

func stateManagementNextStepItem(title, description, href string) Widget {
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

func stateManagementNavigationButtons(previousHref, nextHref string) Widget {
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
				).BackgroundColor("#6B7280").Width(AllBP(120)),
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
				).BackgroundColor("#2B799B").Width(AllBP(120)),
			).Href(nextHref),
		},
	).Flex(1)
}

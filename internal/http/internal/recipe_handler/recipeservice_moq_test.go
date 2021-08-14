// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package recipe_handler

import (
	"github.com/quii/go-http-reference-impl/domain"
	"sync"
)

// Ensure, that RecipeServiceMock does implement RecipeService.
// If this is not the case, regenerate this file with moq.
var _ RecipeService = &RecipeServiceMock{}

// RecipeServiceMock is a mock implementation of RecipeService.
//
// 	func TestSomethingThatUsesRecipeService(t *testing.T) {
//
// 		// make and configure a mocked RecipeService
// 		mockedRecipeService := &RecipeServiceMock{
// 			GetRecipeFunc: func(id string) (domain.Recipe, error) {
// 				panic("mock out the GetRecipe method")
// 			},
// 			StoreRecipeFunc: func(recipe domain.Recipe) (string, error) {
// 				panic("mock out the StoreRecipe method")
// 			},
// 		}
//
// 		// use mockedRecipeService in code that requires RecipeService
// 		// and then make assertions.
//
// 	}
type RecipeServiceMock struct {
	// GetRecipeFunc mocks the GetRecipe method.
	GetRecipeFunc func(id string) (domain.Recipe, error)

	// StoreRecipeFunc mocks the StoreRecipe method.
	StoreRecipeFunc func(recipe domain.Recipe) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetRecipe holds details about calls to the GetRecipe method.
		GetRecipe []struct {
			// ID is the id argument value.
			ID string
		}
		// StoreRecipe holds details about calls to the StoreRecipe method.
		StoreRecipe []struct {
			// Recipe is the recipe argument value.
			Recipe domain.Recipe
		}
	}
	lockGetRecipe   sync.RWMutex
	lockStoreRecipe sync.RWMutex
}

// GetRecipe calls GetRecipeFunc.
func (mock *RecipeServiceMock) GetRecipe(id string) (domain.Recipe, error) {
	if mock.GetRecipeFunc == nil {
		panic("RecipeServiceMock.GetRecipeFunc: method is nil but RecipeService.GetRecipe was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetRecipe.Lock()
	mock.calls.GetRecipe = append(mock.calls.GetRecipe, callInfo)
	mock.lockGetRecipe.Unlock()
	return mock.GetRecipeFunc(id)
}

// GetRecipeCalls gets all the calls that were made to GetRecipe.
// Check the length with:
//     len(mockedRecipeService.GetRecipeCalls())
func (mock *RecipeServiceMock) GetRecipeCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetRecipe.RLock()
	calls = mock.calls.GetRecipe
	mock.lockGetRecipe.RUnlock()
	return calls
}

// StoreRecipe calls StoreRecipeFunc.
func (mock *RecipeServiceMock) StoreRecipe(recipe domain.Recipe) (string, error) {
	if mock.StoreRecipeFunc == nil {
		panic("RecipeServiceMock.StoreRecipeFunc: method is nil but RecipeService.StoreRecipe was just called")
	}
	callInfo := struct {
		Recipe domain.Recipe
	}{
		Recipe: recipe,
	}
	mock.lockStoreRecipe.Lock()
	mock.calls.StoreRecipe = append(mock.calls.StoreRecipe, callInfo)
	mock.lockStoreRecipe.Unlock()
	return mock.StoreRecipeFunc(recipe)
}

// StoreRecipeCalls gets all the calls that were made to StoreRecipe.
// Check the length with:
//     len(mockedRecipeService.StoreRecipeCalls())
func (mock *RecipeServiceMock) StoreRecipeCalls() []struct {
	Recipe domain.Recipe
} {
	var calls []struct {
		Recipe domain.Recipe
	}
	mock.lockStoreRecipe.RLock()
	calls = mock.calls.StoreRecipe
	mock.lockStoreRecipe.RUnlock()
	return calls
}

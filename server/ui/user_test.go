package ui

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/staumann/caluclation/model"
	"net/http"
	"net/http/httptest"
)

type userMock struct {
	getUserHandler  func() []*model.User
	saveUserHandler func(*model.User) error
}

func (u *userMock) GetUserByID(int64) *model.User   { return nil }
func (u *userMock) GetUsers() []*model.User         { return u.getUserHandler() }
func (u *userMock) SaveUser(user *model.User) error { return u.saveUserHandler(user) }
func (u *userMock) UpdateUser(*model.User) error    { return nil }
func (u *userMock) DeleteUserByID(int64) error      { return nil }

var _ = Describe("user ui go", func() {
	BeforeSuite(func() {
		ParseTemplates("../../frontend/html")
	})
	Context("http", func() {
		It("should render the new user template", func() {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/users/new", nil)

			NewUserHandler(recorder, request)
			result := recorder.Result()

			Expect(result.StatusCode).To(Equal(http.StatusOK))
			Expect(recorder.Body.String()).To(ContainSubstring("<head data-test=\"new-user-page\">"))
		})
		It("should render the new user template", func() {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/users", nil)
			userRepository = &userMock{
				getUserHandler: func() []*model.User {
					return []*model.User{
						{
							FirstName: "Peter",
							LastName:  "Griffin",
						},
						{
							FirstName: "Ed",
							LastName:  "Mercer",
						},
					}
				},
			}

			UserHandler(recorder, request)
			result := recorder.Result()

			Expect(result.StatusCode).To(Equal(http.StatusOK))
			Expect(recorder.Body.String()).To(ContainSubstring("<head data-test-id=\"user-overview\">"))
			Expect(recorder.Body.String()).To(ContainSubstring("Griffin, Peter"))
			Expect(recorder.Body.String()).To(ContainSubstring("Mercer, Ed"))
		})
		It("should save a passed user to the database", func() {
			saveWasCalled := false

			userRepository = &userMock{saveUserHandler: func(user *model.User) error {
				Expect(user.LastName).To(Equal("Müller"))
				Expect(user.FirstName).To(Equal("Frank"))
				Expect(user.Image).To(Equal("https://pert.der/image.png"))
				Expect(user.Password).To(Equal("secure112"))
				saveWasCalled = true
				return nil
			}}

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/users/create", nil)
			request.Form = make(map[string][]string)
			request.Form.Set("lastName", "Müller")
			request.Form.Set("firstName", "Frank")
			request.Form.Set("image", "https://pert.der/image.png")
			request.Form.Set("password", "secure112")

			CreateUserHandler(recorder, request)

			Expect(saveWasCalled).To(BeTrue())
			Expect(recorder.Result().StatusCode).To(Equal(http.StatusMovedPermanently))
		})
	})
})

package ui

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/staumann/caluclation/model"
	"net/http"
	"net/http/httptest"
)

type userMock struct {
	getUserHandler func() []*model.User
}

func (u *userMock) GetUserByID(int64) *model.User { return nil }
func (u *userMock) GetUsers() []*model.User       { return u.getUserHandler() }
func (u *userMock) SaveUser(*model.User) error    { return nil }
func (u *userMock) UpdateUser(*model.User) error  { return nil }
func (u *userMock) DeleteUserByID(int64) error    { return nil }

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
	})
})

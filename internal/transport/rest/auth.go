package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vslitvinov/entebot/internal/models"
	"github.com/vslitvinov/entebot/internal/service"
)

type AuthServise interface {
	EmailSingIn(ctx context.Context, email, password string, d service.Device)  (models.Session, error)
	CheckSession(ctx context.Context, sid string) (models.Session, error) 
	SingUp(ctx context.Context, ma models.Accounty) (string,error)
	LogOut(ctx context.Context, sid string) error
}


//Handler
type Auth struct {
	service AuthServise
}


func NewAuth (as service.AccountService, ss service.SessionService) *Auth{
	return &Auth{
		service: service.NewAuth(as,ss),
	}
}


func (h *Auth) checkCookieSession(r *http.Request) (bool, error) {

	// токен сеанса из cookies
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return false, fmt.Errorf("heandler.getCookie %w", err)

		}
		return false, fmt.Errorf("heandler.getCookie %w", err)

	}

	// получаем имя пользователя из кеша сеанса

	_, err = h.service.CheckSession(context.TODO(),c.Value)
	if err != nil {
		return false, err 
	}

	// todo
	// if s.ExpiresAt.Before(time.Now()) {

	// }

	// истекло ли время сессии 
	// if userSession.(sessionUser).isExpired() {
	// 	us.cacheSession.Delete(c.Value)
	// 	return false
	// }

	return true, nil
}

func (h *Auth) SingIn(w http.ResponseWriter, r *http.Request){



	// if r.Method != "POST" {
	// 	return 
	// }

	// // ok, err := h.checkCookieSession(r)

	// r.ParseForm()                    
    // email := r.Form.Get("email")
    // password := r.Form.Get("password")

    // d := service.Device{
    // 	UserAgent: r.Header.Get("User-Agent"),
    // 	IP:        r.RemoteAddr,
    // }

	// s, err := h.service.EmailSingIn(r.Context(),email,password,d)
	// if err != nil {
	// 	log.Println(err)
	// }



}


func (h *Auth) SingUp(w http.ResponseWriter, r *http.Request){}
func (h *Auth) LogOut(w http.ResponseWriter, r *http.Request){}
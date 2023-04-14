import { loginSuccess, logout } from "../../utils/reducers/auth";
import { baseURL } from "../API";
//SIGN UP
function validateName(name) {
  const regex = /^[A-Z][a-z]*$/;
  return regex.test(name);
}

function validateSurname(surname) {
  const regex = /^[A-Z][a-z]*(-[A-Z][a-z]*)?$/;
  return regex.test(surname);
}

function validateEmail(email) {
  const regex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  return regex.test(email);
}

function validatePassword(password) {
  if (password.length < 8) {
    return false;
  }

  const hasUpper = /[A-Z]/.test(password);
  const hasLower = /[a-z]/.test(password);
  const hasDigit = /\d/.test(password);

  return hasUpper && hasLower && hasDigit;
}

const validateData = (mail, name, surname, password, setStatus) => {
    if (!validateEmail(mail)) {
        setStatus('Invalid Email.')
        return false
    }
    if (!validateName(name)){
        setStatus('Invalid name format')
        return false
    }
    if (!validateSurname(surname)){
      setStatus('Invalid surname format')
        return false
    }
    if (!validatePassword(password)){
        setStatus('Invalid Password.')
        return false
    }
    setStatus('')
    return true
}
export const handleSignUp = async (e, mail, name, surname, password, setStatus, setError, navigate) => {
    e.preventDefault();
    
    // Data validation
    if (validateData(mail, name, surname, password, setStatus)){
        try{
            await fetch(`/signup`, 
            {
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify({
                    email: mail,
                    name: name,
                    surname: surname,
                    password: password
                })
            }).then(async (r) => {
                if (!r.ok){
                    setStatus(r.statusText)
                    console.log("Error");
                } else if (r.ok){
                  setError(null);
                  setStatus('');
                  navigate("/signin", {
                      state: {
                        success: true,
                      },
                    });
                }
            })    
        } catch(error){
            console.log(error);
            setError(error);
        }
      } 
    
}

//SIGN IN
export const handleLogin = async (e, email, password, dispatch, navigate, setStatus, setError) => {
    e.preventDefault();
    if (validateEmail(email) && validatePassword(password)) {
      try {
        await fetch(`/login`, {
          headers: {
            Accept: 'application/json',
            'Content-type': 'application/json',
            Credentials: 'include',
          },
          method: 'POST',
          credentials: 'include',
  
          body: JSON.stringify({
            email: email,
            password: password,
          }),
        })
          .then(async (r) => {
            if (!r.ok) {
               throw r.statusText;
            } else if (r.ok) {
              // Parse the response body as JSON
              const data = await r.json();
  
              // Store the JWT token in local storage
              localStorage.setItem('token', data.token);
              localStorage.setItem('user_name', data.user_name);
              
              // Store Auth data in Redux
              dispatch(loginSuccess({ username: data.user_name }))              
              setStatus('');
              navigate('/');
            }
          });
      } catch (error) {
        console.log(error);
        setStatus(error);
      }
    } else if (!validateEmail(email)){
      setStatus('Invalid Email.')

    } else if (!validatePassword(password)){
      setStatus('Invalid Password.')
    }
  };

//HANDLE LOG OUT
export const signOutHandler = (dispatch, navigate) =>{
  // Remove the token from local storage
  localStorage.removeItem('token');
  // Dispatch the logout action to update the Redux state
  dispatch(logout());

  navigate("/signin");  
}
export const signOutOnError = async () =>{ // Very unlikely to happen
    await fetch(`${baseURL}/auth/log_out`, {
        headers: {
            'Accept': 'application/json',
            'Credentials': 'include',
        },
        method: "GET",
        credentials: "include",
        })
}
      
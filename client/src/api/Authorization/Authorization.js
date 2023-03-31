import { loginSuccess, logout } from "../../utils/reducers/auth";
import { baseURL } from "../API";
//SIGN UP
export const handleSignUp = async (e, mail, name, surname, password, setStatus, setError, navigate) => {
    e.preventDefault();

    // const validateData = (email, username, password) => {
    //     const patternUsername = /^[a-zA-Z][a-zA-Z0-9._-]{3,15}$/;
    //     const patternMail = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{3,29}$/;
    //     const patternPassword = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&_])[A-Za-z\d@$!%*?&_]{7,}$/;
    //     if (!patternUsername.test(username)) {
    //         setStatus('Invalid Username. It must be 3 to 15 characters long and can\'t start with a number')
    //         return false
    //     }
    //     if (!patternMail.test(email)){
    //         setStatus('Invalid email format')
    //         return false
    //     }
    //     if (!patternPassword.test(password)){
    //         setStatus('Invalid Password. A valid password has at least one lower and one uppercase letter, one digit, one special character and must be at least 7 characters long')
    //         return false
    //     }
    //     return true
    // }
    //Data validation
    // if (validateData(mail, username, password)){
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
                    console.log("Error");
                }
                setError(null);
                setStatus('');
                navigate("/signin", {
                    state: {
                      success: true,
                    },
                  });
            })    
        } catch(error){
            console.log(error);
            setError(error);
        }
    
}

//SIGN IN
export const handleLogin = async (e, email, password, dispatch, navigate, setStatus, setError) => {
    e.preventDefault(); 
    if (email && password){
      try {
        await fetch(`/login`, 
        {
            headers: {
                'Accept': 'text/plain',
                'Content-type': 'text/plain',
                'Credentials': 'include'
            },  
            method: 'POST',
            credentials: 'include',
    
            body: JSON.stringify({
                email: email,
                password: password
            }),
        }).then(async (r) => {
            if (!r.ok){
                navigate(0);
            } else if (r.ok) {
                setError(null);
                setStatus('');
                navigate('/')
            }
          }) 
      } catch (error) {
            console.log(error);
            setError(error)
      }  
    } else {
        setStatus('Missing authorization data')
    }
}

//HANDLE LOG OUT
export const signOutHandler = async (dispatch, navigate) =>{
    const response = await fetch(`${baseURL}/auth/log_out`, {
        headers: {
            'Accept': 'application/json',
            'Credentials': 'include',
        },
        method: "GET",
        credentials: "include",
        })
        
        dispatch(logout());
        if (!response.ok){
            const error = new Error('Error fetching token data for the Sign Out');
            error.status = response.status;
            navigate("/signin", {state: {error}})
        }
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
      
import { redirect } from "react-router";


export const homeLoader = () => {
    
    const token = localStorage.getItem('token');

    if (token) {
        return (redirect("/statistics"))
    } else {
        return (redirect("/signin"))
    }
}
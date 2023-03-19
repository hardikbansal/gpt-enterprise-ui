import React, { useState } from 'react';
import {useNavigate, Routes} from 'react-router-dom'
import { googleLogout, useGoogleLogin } from '@react-oauth/google';
import axios from 'axios';

function Login() {
    const redirect = useNavigate();

    const login = useGoogleLogin({
        onSuccess: (codeResponse) => loginUser(codeResponse),
        onError: (error) => console.log('Login Failed:', error)
    });

    const testLogin = () =>{
        console.log(localStorage.setItem('access_token',JSON.stringify({'token':'sample_token'})));
        console.log(JSON.parse(localStorage.getItem('access_token')));
    }

    const loginUser = (user) => {
        localStorage.setItem("sso_token", user.access_token);
        console.log(user);
        const headers = {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': 'Bearer ' + user.access_token
        }          
        axios
            .post('http://localhost:3000/api/accesstoken', {},{
                headers: headers 
            })
            .then((res) => {
                console.log(res.data);
                localStorage.setItem("access_token", res.data);
                redirect("/chat");
            })
            .catch((err) => {
                console.log(err);
                logOut();
            });
    };
        

    // log out function to log the user out of google and set the profile array to null
    const logOut = () => {
        googleLogout();
        localStorage.removeItem("sso_token");
        localStorage.removeItem("access_token");
    };

    return (
        <div>
            <button onClick={() => testLogin()}>Sign in with Google 🚀 </button>
        </div>
    );
}
export default Login;
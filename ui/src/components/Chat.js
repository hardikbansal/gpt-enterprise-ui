import React, { useState } from 'react';
import {useNavigate, Routes} from 'react-router-dom'
import { googleLogout, useGoogleLogin } from '@react-oauth/google';
import axios from 'axios';

function Chat() {
    const redirect = useNavigate();

    const login = useGoogleLogin({
        onSuccess: (codeResponse) => loginUser(codeResponse),
        onError: (error) => console.log('Login Failed:', error)
    });

    const loginUser = (user) => {
        localStorage.setItem("sso_token", user.access_token);
        console.log(user);
        axios
            .get(`http://localhost:8080/api/accesstoken`, {
                headers: {
                    Authorization: `Bearer ${user.access_token}`,
                    Accept: 'application/json'
                }
            })
            .then((res) => {
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
            <button onClick={() => login()}>Sign in with Google 🚀 on chat page </button>
        </div>
    );
}
export default Chat;
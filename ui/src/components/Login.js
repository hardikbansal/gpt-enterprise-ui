import React, { useState } from 'react';
import {useNavigate, Routes} from 'react-router-dom'
import { googleLogout, useGoogleLogin } from '@react-oauth/google';
import axios from 'axios';

function Login() {
    const ACCESS_TOKEN =  process.env.REACT_APP_API_ENDPOINT+"/api/accesstoken"
    const redirect = useNavigate();
    const isDebug = process.env.REACT_APP_DEBUG==="true";

    const login = useGoogleLogin({
        onSuccess: (codeResponse) => loginUser(codeResponse),
        onError: (error) => console.log('Login Failed:', error)
    });

    const testLogin = () =>{
        loginUser({"access_token":"sample_token"})
    }

    const loginUser = (user) => {
        localStorage.setItem("sso_token", user.access_token);
        console.log("Google user: ",user);
        const headers = {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': 'Bearer ' + user.access_token
        }          
        axios
            .get(ACCESS_TOKEN,{
                headers: headers 
            })
            .then((res) => {
                console.log("Login token from server: ", res.data);
                localStorage.setItem("user_token", res.data.token);
                redirect("/");
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
        localStorage.removeItem("user_token");
        redirect('/login');
    };

    return (
        <div className="h-full flex flex-col items-center justify-center">
            <div className="flex flex-col items-center justify-center">
                <button onClick={() => isDebug? testLogin(): login()} aria-label="Continue with google" role="button"
                        className="focus:outline-none focus:ring-2 focus:ring-offset-1 focus:ring-gray-700 py-3.5 px-4 border rounded-lg border-gray-700 flex items-center w-full mt-10">
                    <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/sign_in-svg2.svg" alt="google"/>
                    <p className="text-base font-medium ml-4 text-gray-700">Continue with Google</p>
                </button>
            </div>

        </div>
    );
}
export default Login;

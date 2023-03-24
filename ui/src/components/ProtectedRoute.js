import React, { useEffect, useState } from "react";
import { Route, useNavigate } from "react-router-dom";
import axios from "axios";
import {googleLogout} from "@react-oauth/google";
const ProtectedRoute = (props) => {
	console.log(process.env.REACT_APP_API_ENDPOINT)
    const USER_DETAILS = process.env.REACT_APP_API_ENDPOINT+"/api/user"
    const navigate = useNavigate();
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const checkUserToken = () => {
        const userToken = localStorage.getItem('user_token');
        console.log("User Token: ", userToken);
        if (!userToken || userToken === 'undefined') {
            setIsLoggedIn(false);
            return navigate('/login');
        }
        const headers = {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': 'Bearer ' + userToken
        }
        axios.get(USER_DETAILS,
            {headers: headers}
        )
            .then(resp => {
                setIsLoggedIn(true);
            })
            .catch(error => {
                if (error.response.status === 401){
                   logOut();
                }
                console.log(error);
            })

    }

    // log out function to log the user out of google and set the profile array to null
    const logOut = () => {
        googleLogout();
        localStorage.removeItem("sso_token");
        localStorage.removeItem("user_token");
        navigate('/login');
    };

    useEffect(() => {
        checkUserToken();
    }, [isLoggedIn]);
    return (
        <React.Fragment>
            {
                isLoggedIn ? props.children : null
            }
        </React.Fragment>
    );
}
export default ProtectedRoute;

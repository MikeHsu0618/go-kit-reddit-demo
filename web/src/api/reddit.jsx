import axios from "axios";
import {toast} from "react-toastify";

export const CreatePost = async(title, content, token, userId) => {
    try {
        const res = await axios({
            method: 'POST',
            url: "http://localhost:8381/create-post",
            headers: {
                Authorization: `Bearer ${token}`
            },
            data: {
                title: title,
                content: content,
                user_id: userId,
            }
        })
        return res.data
    } catch (e) {
        toast.error(e.response.data.error)
        return ''
    }
}

export const ListPost = async(token) => {
    try {
        const res = await axios({
            method: 'GET',
            url: "http://localhost:8381/list-post",
            headers: {
                Authorization: `Bearer ${token}`
            }
        })
        return res.data
    } catch (e) {
        toast.error(e.response.data.error)
    }
}

export const userLogin = async (username, password) => {
    let token = ''
    try {
        const res = await axios({
            method: 'POST',
            url: "http://localhost:8381/login",
            data: {
                username: username,
                password: password
            }
        })
        token = res.data
    } catch (e) {
        toast.error(e.response.data.error)
    }
    return token

}

export const userRegister = async (username, password) => {
    try {
        const res = await axios({
            method: 'POST',
            url: "http://localhost:8381/register",
            data: {
                username: username,
                password: password
            }
        })
        return res.data
    } catch (e) {
        toast.error(e.response.data.error)
    }
}
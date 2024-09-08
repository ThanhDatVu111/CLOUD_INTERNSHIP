import React, { useEffect, useState } from "react";

export function useUsers() {
    const [userList, setUserList] = useState([]);

    useEffect(() => {
        fetch("http://localhost:5174/users")
            .then(res => res.json())
            .then(data => {
                setUserList(data);
            })
            .catch(err => console.error("Error getting users", err));
    }, []);

    return userList;
}

export default useUsers;
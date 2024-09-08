import React, { useEffect, useState } from "react";

export function useProjects() {
    const [projectList, setProjectList] = useState([]);

    useEffect(() => {
        fetch("http://localhost:5174/projects")
            .then(res => res.json())
            .then(data => {
                setProjectList(data);
            })
            .catch(err => console.error("Error getting projects", err));
    }, []);

    return projectList;
}

export default useProjects;
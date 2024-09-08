import React, { useEffect } from "react";

function FadeIn({ time, children }) { //pass a number in ms for interval between fade in
    useEffect(() => {
        const elements = document.querySelectorAll('.fadeIn');
        elements.forEach((el, index) => {
            setTimeout(() => {
                el.classList.add('fadeInAnimation');
                el.classList.remove('fadeIn');
            }, index * time); //interval
        });
    }, []);

    return <>{children}</>;
}

export default FadeIn;
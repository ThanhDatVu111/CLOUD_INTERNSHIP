import { useState, useEffect } from 'react';

export function useScrollPosition() {
    const [scrollPosition, setScrollPosition] = useState(0);

    useEffect(() => {
        const handleScroll = () => {
            setScrollPosition(window.scrollY);
        };

        window.addEventListener('scroll', handleScroll);

        // Cleanup function to remove the event listener when the component unmounts
        return () => window.removeEventListener('scroll', handleScroll);
    }, []);

    return scrollPosition;
}
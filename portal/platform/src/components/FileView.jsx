import '../ctStyle.css';

import fileImage from "../images/icons/grayScheme/CTGrayDataIcon.png";

function FileView({type, src, title, description}) {

    //declare supported file types
    const imageFiles = ['jpg', 'jpeg', 'png', 'gif', 'svg'];
    const videoFiles = ['mp4', 'webm'];

    if (imageFiles.includes(type)) {
        type = 'image';
    } else if (videoFiles.includes(type)) {
        type = 'video';
    }

    return (
        <div className='fileView'>
            {type === 'image' ? (
                <img src={src} alt={title} />
            ) : type === 'video' ? (
                <video src={src} controls />
            ) : (
                <img src={fileImage} className='notFull'></img>
            )}

            <div className='cover'>
                <h3>{title}</h3>
                <p>{description}</p>
            </div>
        </div>
    )
}

export default FileView;
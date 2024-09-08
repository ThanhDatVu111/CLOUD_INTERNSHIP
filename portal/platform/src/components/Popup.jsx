import '../ctStyle.css';

function popup({type, header, message, preMessage, preTime}) {
    
    //delete popup
    function hide() {
        //delete the popup element
        let popup = document.querySelector('.popup');
        popup = popup.parentElement;
        popup.remove();
    }

    return (
        <div className='popup fadeInFast'>
            <div className='popupBox'>
                {type == 'addEvent' ? ( //when validating, strings should never contain backslashes
                    <div>
                        <h2>{header}</h2>
                        <p>Enter event name</p>
                        <input type="text" id="userInput" className="smallBottom" />
                        <p>Enter event time (optional)</p>
                        <input type="time" id="timeInput" className="smallBottom" />
                        <div className='flexCenter'>
                            <button onClick={() => hide()} id="false">Cancel</button>
                            <button onClick={() => hide()} id="true">Okay</button>
                        </div>
                    </div>
                ) : type == 'editEvent' ? (
                    <div>
                        <h2>{header}</h2>
                        <p>Edit event name</p>
                        <input type="text" id="userInput" className="smallBottom" defaultValue={preMessage} />
                        <p>Edit event time (optional)</p>
                        <input type="time" id="timeInput" className="smallBottom" defaultValue={preTime} />
                        <div className='flexCenter'>
                            <button onClick={() => hide()} id="false">Cancel</button>
                            <button onClick={() => hide()} id="true">Okay</button>
                        </div>
                    </div>
                ) : type == "boolean" ? ( 
                    <div>
                        <h2>{header}</h2>
                        <p>{message}</p>
                        <div className='flexCenter'>
                            <button onClick={() => hide()} id="false">Cancel</button>
                            <button onClick={() => hide()} id="true">Okay</button>
                        </div>
                    </div>
                ) : (
                    <div>
                        <p>Something went wrong. Try again</p>
                        <div className='flexCenter'>
                            <button onClick={() => hide()}>Okay</button>
                        </div>
                    </div>
                )
                }
            </div>
        </div>
    )
}

export default popup;
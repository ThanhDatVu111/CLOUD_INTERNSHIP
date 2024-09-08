import React, {useEffect, useState} from "react";
import '../ctStyle.css';

import * as dataCaller from '../scripts/dataCaller';

function Profile({user, size}) {

    return (
        <div className='profile'>
            <img src={dataCaller.getUser(user).pfp} className="pfp" width={size} height={size} alt={`${dataCaller.getName(user)}'s Profile Picture`}></img>
            <span className='pfpText'>{dataCaller.getName(user)}</span>
        </div>
    )
}

export default Profile;
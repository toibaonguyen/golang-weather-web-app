const inputBox = document.querySelector('.input-box');
const searchBtn = document.getElementById('searchBtn');
const weather_img = document.querySelector('.weather-img');
const temperature = document.querySelector('.temperature');
const description = document.querySelector('.description');
const humidity = document.getElementById('humidity');
const wind_speed = document.getElementById('wind-speed');
const location_not_found = document.querySelector('.location-not-found');
const weather_body = document.querySelector('.weather-body');


const BASE_URL="http://localhost:8080"


async function checkWeather(city){


    const url = `${BASE_URL}?city=${city}`;

    const weather_data = await fetch(`${url}`).then(response => response.json());
    console.log("weather_data:",weather_data)

    if(weather_data.statusCode!="OK"){
        location_not_found.style.display = "flex";
        weather_body.style.display = "none";
        return;
    }

    location_not_found.style.display = "none";
    weather_body.style.display = "flex";
    temperature.innerHTML = `${Math.round(weather_data.data.tempotary - 273.15)}°C`;
    description.innerHTML = `${weather_data.data.description}`;

    humidity.innerHTML = `${weather_data.data.humidity}%`;
    wind_speed.innerHTML = `${weather_data.data.windSpeed}Km/H`;


    switch(weather_data.data.main){
        case 'Clouds':
            weather_img.src = "/assets/cloud.png";
            break;
        case 'Clear':
            weather_img.src = "/assets/clear.png";
            break;
        case 'Rain':
            weather_img.src = "/assets/rain.png";
            break;
        case 'Mist':
            weather_img.src = "/assets/mist.png";
            break;
        case 'Snow':
            weather_img.src = "/assets/snow.png";
            break;
        default:
            weather_img.src = "/assets/clear.png";

    }

}


searchBtn.addEventListener('click', ()=>{
    checkWeather(inputBox.value);
});
import React, { useEffect, useState } from 'react';
import { getProfile } from '../api';
import Chart from './Chart'; // Assuming Chart component path

const Dashboard = () => {
  const [weatherData, setWeatherData] = useState(null);

  useEffect(() => {
    const fetchWeatherData = async () => {
      const token = localStorage.getItem('token');
      try {
        const { data } = await getProfile(token);
        setWeatherData(data.weather_data);
      } catch (error) {
        console.error('Error fetching profile:', error);
      }
    };

    fetchWeatherData();
  }, []);

  if (!weatherData) {
    return <div>Loading...</div>;
  }

  const { location, years } = weatherData;
  // console.log("years dash", years)
  return (
    <div>
      <h2>Dashboard</h2>
      <h3>Weather Forecast for {location.name}, {location.region}, {location.country}</h3>
      {/* <pre>{JSON.stringify(weatherData, null, 2)}</pre> */}
      {/* Render your Chart component */}
      
      <Chart weatherData={years} />
    </div>
  );
};

export default Dashboard;



// import React, { useEffect, useState } from 'react';
// import { getProfile } from '../api';

// const Dashboard = () => {
//   const [weatherData, setWeatherData] = useState(null);

//   useEffect(() => {
//     const fetchWeatherData = async () => {
//       const token = localStorage.getItem('token');
//       try {
//         const { data } = await getProfile(token);
//         setWeatherData(data.weather_data);
//       } catch (error) {
//         console.error('Error fetching profile:', error);
//       }
//     };

//     fetchWeatherData();
//   }, []);

//   if (!weatherData) {
//     return <div>Loading...</div>;
//   }

//   return (
//     <div>
//       <h2>Dashboard</h2>
//       {/* Render your weather data here */}
//       <pre>{JSON.stringify(weatherData, null, 2)}</pre>
//     </div>
//   );
// };

// export default Dashboard;

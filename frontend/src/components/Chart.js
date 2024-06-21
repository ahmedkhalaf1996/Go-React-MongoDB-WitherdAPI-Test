
import React from 'react';
import { Line } from 'react-chartjs-2';
import { 
  Chart as ChartJS, 
  CategoryScale, 
  LinearScale, 
  PointElement, 
  LineElement, 
  Title,
  Tooltip,
  Legend
 } from 'chart.js'

ChartJS.register(  CategoryScale, 
  LinearScale, 
  PointElement, 
  LineElement, 
  Title,
  Tooltip,
  Legend)

const Chart = ({ weatherData }) => {
  const labels = weatherData.forecastday.map((day) => day.date);
  const temperatures = weatherData.forecastday.map((day) => day.avgtemp_c);

  console.log("lables", labels)
  console.log("tempratures", temperatures)
  const data = {
    labels,
    datasets: [
      {
        label: 'Average Temperature (°C)',
        data: temperatures,
        fill: false,
        backgroundColor: 'rgb(75, 192, 192)',
        borderColor: 'rgba(75, 192, 192, 0.2)',
      },
    ],
  };

  const chartOptions = {}
  return <Line data={data} options={chartOptions} />;
};

export default Chart;


import { Chart } from "react-google-charts";


export const data = [
    ["Task", "Hours per Day"],
    ["Work", 11],
    ["Eat", 2],
    ["Commute", 2],
    ["Watch TV", 2],
    ["Sleep", 7], // CSS-style declaration
  ];
  
  export const options = {
    title: "My Spendings",
    pieHole: 0.6,
    is3D: false,
    backgroundColor: "secondary.dark"
  };
  
  const DonutChart = () => {
    return (
      <Chart
        chartType="PieChart"
        width="100%"
        height="600px"
        data={data}
        options={options}
      />
    );
  }

  export default DonutChart;
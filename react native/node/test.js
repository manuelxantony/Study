// console.log("hello");

// let arr = [];

// const setArr = (arr, val) => {
//   return [...arr, val];
// };

// arr = setArr(arr, { tt: "tt" });
// arr = setArr(arr, { rr: "rr" });

// let t = {
//   airDatas: {
//     aqi: 1,
//     cityName: '"Hyde Park"',
//     humidity: 88,
//     temparature: 285.55,
//     uvi: 0.06,
//     windSpeed: 9.77,
//   },
//   pollutans: {
//     co: 211.95,
//     no2: 9.6,
//     o3: 57.22,
//     pm10: 2.44,
//     pm2_5: 2.34,
//     so2: 3.93,
//   },
// };
// console.log("----------", t?.airDatas);

// // arr = setArr(arr, { t });
// // arr = setArr(arr, { t });

// // for (i in arr) {
// //   console.log("arry:", arr[i]);
// // }

// let d = new Date();
// console.log(d.getMilliseconds());

// import { categoryHeaderdata } from "./cat";

// for (const c of categoryHeaderdata) {
//   console.log(c);
// }

class ItemModal {
  constructor(itemImage, imageName, imageNumber, eth) {
    this.itemImage = itemImage;
    this.imageName = imageName;
    this.imageNumber = imageNumber;
    this.eth = eth;
  }
}

const ITEMS = [
  new ItemModal("../assets/pickles/1.png", "Fucking Pickle", "674"),
  new ItemModal("../assets/pickles/2.png", "Fucking Pickle", "645"),
  new ItemModal("../assets/pickles/3.png", "Fucking Pickle", "734"),
  new ItemModal("../assets/pickles/4.png", "Fucking Pickle", "428"),
  new ItemModal("../assets/pickles/5.png", "Fucking Pickle", "049"),
  new ItemModal("../assets/pickles/6.png", "Fucking Pickle", "556"),
];

const render = (item) => {
  return item.imageNumber + "wow";
};

console.log(ITEMS.map((item) => render(item)));

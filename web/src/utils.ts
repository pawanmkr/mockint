const weekday = ["Sun","Mon","Tue","Wed","Thu","Fri","Sun"]
const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

export function humanTime(ts: string) {
  const t = new Date(ts)
  const wkd = `${weekday[t.getDay()]}`
  const date = `${t.getDate()} ${months[t.getMonth()]}`;
  
  let ampm = "AM";
  let hrs: string | number = t.getHours();
  if (hrs > 12) {
    ampm = "PM"
    hrs = t.getHours() - 12;
    if (hrs < 10) {
      hrs = `0${hrs}`
    }
  }

  let mns: string | number = t.getMinutes();
  if (mns < 10) {
    mns = `0${mns}`
  }

  const time = `${hrs}:${mns} ${ampm}`

  return `${wkd}, ${date} ${time}`;
}
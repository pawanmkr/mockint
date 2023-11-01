type Guest = {
  name: string,
  whatsapp: string,
  note: string
}

export type TypeInterview = {
  id: string,
  name: string,
  time: string,
  duration: number,
  skills: string,
  guestType: string,
  difficulty: string,
  note: string,
  booked: boolean,
  meetingCode: string,
  joinUrl: string,
  guest: [Guest]
}
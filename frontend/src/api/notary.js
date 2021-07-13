import fetch from './fetch'
export function getNotary(data) {
  return fetch({
    url: "/api/v1/notaries",
    data
  })
}
export function getAllocated(data) {
  return fetch({
    url: "/api/v1/allocated",
    data
  })
}
export function getDeals(data) {
  return fetch({
    url: '/api/v1/deals',
    data
  })
}
export function getGrantedDaily(data) {
  return fetch({
    url: '/api/v1/granted-daily',
    data
  })
}
export function getProportionOfAllowance(data) {
  return fetch({
    url: '/api/v1/proportion-of-allowance',
    data
  })
}
export function getProportionOfAllowanceByLocation(data) {
  return fetch({
    url: '/api/v1/proportion-of-allowance-by-location',
    data
  })
}

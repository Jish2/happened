# Cloud Run outputs
output "service_name" {
  value = google_cloud_run_v2_service.api.name
  description = "Name of the created service"
}

output "revision" {
  value = google_cloud_run_v2_service.api.latest_created_revision
  description = "Deployed revision for the service"
}

output "service_url" {
  value = google_cloud_run_v2_service.api.uri
  description = "The URL on which the deployed service is available"
}

output "service_id" {
  value = google_cloud_run_v2_service.api.id
  description = "Unique Identifier for the created service"
}

output "service_status" {
  value = google_cloud_run_v2_service.api.traffic_statuses
  description = "Status of the created service"
}

output "service_location" {
  value       = google_cloud_run_v2_service.api.location
  description = "Location in which the Cloud Run service was created"
}
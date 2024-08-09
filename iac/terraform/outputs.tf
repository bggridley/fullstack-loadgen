output "secrets_identity_id" {
    value = azurerm_kubernetes_cluster.k8s.kubelet_identity[0].client_id
}
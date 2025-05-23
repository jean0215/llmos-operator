/*
Copyright 2025 llmos.ai.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=management.llmos.ai
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalRoleList is a list of GlobalRole resources
type GlobalRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []GlobalRole `json:"items"`
}

func NewGlobalRole(namespace, name string, obj GlobalRole) *GlobalRole {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("GlobalRole").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedAddonList is a list of ManagedAddon resources
type ManagedAddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ManagedAddon `json:"items"`
}

func NewManagedAddon(namespace, name string, obj ManagedAddon) *ManagedAddon {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("ManagedAddon").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleTemplateList is a list of RoleTemplate resources
type RoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RoleTemplate `json:"items"`
}

func NewRoleTemplate(namespace, name string, obj RoleTemplate) *RoleTemplate {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("RoleTemplate").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleTemplateBindingList is a list of RoleTemplateBinding resources
type RoleTemplateBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RoleTemplateBinding `json:"items"`
}

func NewRoleTemplateBinding(namespace, name string, obj RoleTemplateBinding) *RoleTemplateBinding {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("RoleTemplateBinding").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SettingList is a list of Setting resources
type SettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Setting `json:"items"`
}

func NewSetting(namespace, name string, obj Setting) *Setting {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Setting").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TokenList is a list of Token resources
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Token `json:"items"`
}

func NewToken(namespace, name string, obj Token) *Token {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Token").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradeList is a list of Upgrade resources
type UpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Upgrade `json:"items"`
}

func NewUpgrade(namespace, name string, obj Upgrade) *Upgrade {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Upgrade").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserList is a list of User resources
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []User `json:"items"`
}

func NewUser(namespace, name string, obj User) *User {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("User").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VersionList is a list of Version resources
type VersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Version `json:"items"`
}

func NewVersion(namespace, name string, obj Version) *Version {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Version").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// generated by go-extpoints -- DO NOT EDIT
package extpoints

import (
	"reflect"
	"sync"
)

var registry = struct {
	sync.Mutex
	extpoints map[string]*extensionPoint
}{
	extpoints: make(map[string]*extensionPoint),
}

type extensionPoint struct {
	sync.Mutex
	iface      reflect.Type
	components map[string]interface{}
}

func newExtensionPoint(iface interface{}) *extensionPoint {
	ep := &extensionPoint{
		iface:      reflect.TypeOf(iface).Elem(),
		components: make(map[string]interface{}),
	}
	registry.Lock()
	defer registry.Unlock()
	registry.extpoints[ep.iface.Name()] = ep
	return ep
}

func (ep *extensionPoint) lookup(name string) (ext interface{}, ok bool) {
	ep.Lock()
	defer ep.Unlock()
	ext, ok = ep.components[name]
	return
}

func (ep *extensionPoint) all() map[string]interface{} {
	ep.Lock()
	defer ep.Unlock()
	registered := make(map[string]interface{})
	for k, v := range ep.components {
		registered[k] = v
	}
	return registered
}

func (ep *extensionPoint) register(component interface{}) bool {
	return ep.registerNamed(component, reflect.TypeOf(component).Elem().Name())
}

func (ep *extensionPoint) registerNamed(component interface{}, name string) bool {
	ep.Lock()
	defer ep.Unlock()
	_, exists := ep.components[name]
	if exists {
		return false
	}
	ep.components[name] = component
	return true
}

func (ep *extensionPoint) unregister(name string) bool {
	ep.Lock()
	defer ep.Unlock()
	_, exists := ep.components[name]
	if !exists {
		return false
	}
	delete(ep.components, name)
	return true
}

func implements(component interface{}) []string {
	var ifaces []string
	for name, ep := range registry.extpoints {
		if reflect.TypeOf(component).Implements(ep.iface) {
			ifaces = append(ifaces, name)
		}
	}
	return ifaces
}

func RegisterNamed(component interface{}, name string) []string {
	registry.Lock()
	defer registry.Unlock()
	var ifaces []string
	for _, iface := range implements(component) {
		if ok := registry.extpoints[iface].registerNamed(component, name); ok {
			ifaces = append(ifaces, iface)
		}
	}
	return ifaces
}

func Register(component interface{}) []string {
	registry.Lock()
	defer registry.Unlock()
	var ifaces []string
	for _, iface := range implements(component) {
		if ok := registry.extpoints[iface].register(component); ok {
			ifaces = append(ifaces, iface)
		}
	}
	return ifaces
}

// LifecycleParticipant

var LifecycleParticipants = &lifecycleParticipantExt{
	newExtensionPoint(new(LifecycleParticipant)),
}

type lifecycleParticipantExt struct {
	*extensionPoint
}

func (ep *lifecycleParticipantExt) Unregister(name string) bool {
	return ep.unregister(name)
}

func (ep *lifecycleParticipantExt) Register(component LifecycleParticipant) bool {
	return ep.register(component)
}

func (ep *lifecycleParticipantExt) RegisterNamed(component LifecycleParticipant, name string) bool {
	return ep.registerNamed(component, name)
}

func (ep *lifecycleParticipantExt) Lookup(name string) (LifecycleParticipant, bool) {
	ext, ok := ep.lookup(name)
	return ext.(LifecycleParticipant), ok
}

func (ep *lifecycleParticipantExt) All() map[string]LifecycleParticipant {
	all := make(map[string]LifecycleParticipant)
	for k, v := range ep.all() {
		all[k] = v.(LifecycleParticipant)
	}
	return all
}

// CommandProvider

var CommandProviders = &commandProviderExt{
	newExtensionPoint(new(CommandProvider)),
}

type commandProviderExt struct {
	*extensionPoint
}

func (ep *commandProviderExt) Unregister(name string) bool {
	return ep.unregister(name)
}

func (ep *commandProviderExt) Register(component CommandProvider) bool {
	return ep.register(component)
}

func (ep *commandProviderExt) RegisterNamed(component CommandProvider, name string) bool {
	return ep.registerNamed(component, name)
}

func (ep *commandProviderExt) Lookup(name string) (CommandProvider, bool) {
	ext, ok := ep.lookup(name)
	return ext.(CommandProvider), ok
}

func (ep *commandProviderExt) All() map[string]CommandProvider {
	all := make(map[string]CommandProvider)
	for k, v := range ep.all() {
		all[k] = v.(CommandProvider)
	}
	return all
}


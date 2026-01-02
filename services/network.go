package services

import (
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

func ScanPort(ip string, port int) bool {
	// JoinHostPort pone automáticamente los [] si detecta que es IPv6
	// y maneja el : para el puerto correctamente.
	address := net.JoinHostPort(ip, fmt.Sprintf("%d", port))

	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func GetUltimateScan() string {
	baseIP := "192.168.1"
	var wg sync.WaitGroup

	// 1. FASE DE DESCUBRIMIENTO (Silenciosa)
	// Esto "despierta" a los dispositivos y llena la tabla ARP de Windows
	for i := 1; i <= 254; i++ {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			// No nos importa el resultado del ping, solo que se intente la conexión
			exec.Command("ping", "-n", "1", "-w", "100", ip).Run()
		}(fmt.Sprintf("%s.%d", baseIP, i))
	}
	wg.Wait()

	// 2. FASE DE FORMATO (Llamamos a tu función anterior)
	// Esta función es la que ya tiene el Regex y las IPs bonitas
	return GetCleanARP()
}

func GetCleanARP() string {
	cmd := exec.Command("arp", "-a")
	output, _ := cmd.CombinedOutput()
	rawStr := string(output)

	// 1. Buscamos todas las líneas que tengan una IP y una MAC
	// Este regex busca: [IP] [Espacios] [MAC] [Espacios] [dinmico/esttico]
	re := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s+([a-fA-F0-9-]{17})\s+(\w+)`)
	matches := re.FindAllStringSubmatch(rawStr, -1)

	var builder strings.Builder
	builder.WriteString("🔍 *Dispositivos Reales Detectados:*\n\n")

	for _, match := range matches {
		ip := match[1]
		mac := match[2]
		tipo := match[3]

		// 2. FILTRO DE SEGURIDAD:
		// Ignoramos IPs de Multicast (224+) y Broadcast (.255)
		if strings.HasPrefix(ip, "224") || strings.HasPrefix(ip, "239") || strings.HasSuffix(ip, ".255") {
			continue
		}

		// 3. Formateamos cada dispositivo
		builder.WriteString(fmt.Sprintf("🖥 `%-15s` | `%s` (%s)\n", ip, mac, tipo))
	}

	if builder.Len() < 40 { // Si solo quedó el título
		return "No se encontraron dispositivos activos en el rango privado."
	}

	return builder.String()
}
func CheckHost(ip string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	// 1. Intentamos PING (Universal para detectar si está encendido)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "-w", "500", ip)
	} else {
		cmd = exec.Command("ping", "-c", "1", "-W", "1", ip)
	}

	if err := cmd.Run(); err != nil {
		return // Si no hay ping, asumimos que no hay nada
	}

	// 2. Si hay ping, investigamos qué es (Fingerprinting)
	tipo := "📱 Celular/Tablet" // Por defecto, si solo hay ping

	// Puertos típicos de PCs o Servidores (Windows, Linux, Mac)
	// 135, 445: Windows | 22: SSH (Linux/Mac) | 5900: VNC
	pcPorts := []int{135, 445, 22, 5900}

	for _, port := range pcPorts {
		address := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
		conn, err := net.DialTimeout("tcp", address, 200*time.Millisecond)
		if err == nil {
			conn.Close()
			tipo = "💻 Computadora/Servidor"
			break
		}
	}

	// 3. Intentamos obtener el nombre (Hostname)
	names, _ := net.LookupAddr(ip)
	hostname := ""
	if len(names) > 0 {
		hostname = fmt.Sprintf(" (%s)", names[0])
	}

	results <- fmt.Sprintf("%s `%s`%s", tipo, ip, hostname)
}

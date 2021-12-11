from scapy.all import *


def checkDNS(name):
    ip = name
    try:
        ans = sr1(IP(dst=ip) / UDP(sport=RandShort(), dport=53) / DNS(rd=1,
                                                                      qd=DNSQR(qname="google.com", qtype="TXT")), verbose=False, timeout=0.5)
        ansTXT = ans[DNSRR][0].rdata
        ans = sr1(IP(dst=ip) / UDP(sport=RandShort(), dport=53) / DNS(rd=1,
                                                                      qd=DNSQR(qname="amazon.com", qtype="A")), verbose=False, timeout=0.5)
        ansA1 = ans[DNSRR][0].rdata
        ans = sr1(IP(dst=ip) / UDP(sport=RandShort(), dport=53) / DNS(rd=1,
                                                                      qd=DNSQR(qname="google.com", qtype="A")), verbose=False, timeout=0.5)
        ansA2 = ans[DNSRR][0].rdata
        if ansTXT == ansA1 and ansA1 == ansA2 and ansA2 != '' and valid_ip(ansA1):
            print("[+] Cs Detected: " + ip + " replied with: " + ansTXT)
    except Exception as e:
        print(e)
        pass


if __name__ == '__main__':
    name = "192.168.62.136"
    checkDNS(name)

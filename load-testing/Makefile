
generate-html:
	rm -r ./html-report/*
	jmeter -g login.jtl -o ./html-report

run-tests:
	HEAP="-Xms1g -Xmx1g -XX:MaxMetaspaceSize=256m" jmeter -n -t test-plan.jmx -l result.jtl -e -o html-report

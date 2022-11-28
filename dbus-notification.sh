app_name="app"
id="10"
icon="error"
summary="summary"
body="body"
actions="[]"
hints="{}"
timeout="5000"

gdbus call --session   \
   --dest org.freedesktop.Notifications \
   --object-path /org/freedesktop/Notifications \
   --method org.freedesktop.Notifications.Notify \
   "${app_name}" "${id}" "${icon}" "${summary}" "${body}" \
   "${actions}" "${hints}" "${timeout}"

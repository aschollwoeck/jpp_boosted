from django import template
from django.template.loader import get_template
register = template.Library()

@register.inclusion_tag('project_overview_list.html')
def project_overview_list(project_list):
    return {'project_list': project_list}